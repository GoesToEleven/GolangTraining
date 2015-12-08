package githubexample

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/delay"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

const redirectURI = "http://localhost:8080/oauth2callback"
const githubAPIURL = "https://api.github.com"

var delayedGetStats *delay.Function

func init() {
	delayedGetStats = delay.Func(
		"my-favorite-dog",
		func(ctx context.Context, accessToken, username string) error {
			api := &GithubAPI{
				ctx:         ctx,
				accessToken: accessToken,
				username:    username,
			}
			since := time.Now().Add(-time.Hour * 24 * 30)
			stats, err := api.getCommitSummaryStats(since)
			if err != nil {
				return err
			}
			log.Infof(ctx, "STATS: %v", stats)

			key := datastore.NewKey(ctx, "Stats", username, 0, nil)
			_, err = datastore.Put(ctx, key, &stats)
			if err != nil {
				return err
			}
			return nil
		},
	)
}

type GithubAPI struct {
	ctx         context.Context
	accessToken string
	username    string
}

type CommitStats struct {
	Additions, Deletions int
}

func NewGithubAPI(ctx context.Context) *GithubAPI {
	return &GithubAPI{
		ctx: ctx,
	}
}

func (api *GithubAPI) getUsername() (string, error) {
	client := urlfetch.Client(api.ctx)
	response, err := client.Get(
		githubAPIURL + "/user?access_token=" + api.accessToken)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var data struct {
		Login string
	}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	return data.Login, nil
}

func (api *GithubAPI) getAccessToken(state, code string) (string, error) {
	values := make(url.Values)
	values.Add("client_id", "767154e6915134caade5")
	values.Add("client_secret", "50648a71510c21bb77e229692fc882dfbe8ea35d")
	values.Add("code", code)
	values.Add("state", state)

	client := urlfetch.Client(api.ctx)
	response, err := client.PostForm("https://github.com/login/oauth/access_token", values)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bs, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	values, err = url.ParseQuery(string(bs))
	if err != nil {
		return "", err
	}
	return values.Get("access_token"), nil
}

func (api *GithubAPI) getCommitSummaryStats(since time.Time) (CommitStats, error) {
	var stats CommitStats

	type Repo struct {
		Organization, Repository string
	}

	var wg sync.WaitGroup
	repositoryChannel := make(chan Repo)
	statsChannel := make(chan CommitStats)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for repo := range repositoryChannel {
				// list all commits for repository: GET /repos/:owner/:repo/commits
				shas, err := api.getUserCommitShas(repo.Organization, repo.Repository, since)
				if err != nil {
					log.Errorf(api.ctx, "Error getting user commit shas: %v", err)
					return
				}
				log.Infof(api.ctx, "REPO: %v, SHAS:%v", repo, shas)
				for _, sha := range shas {
					// get a single commit: GET /repos/:owner/:repo/commits/:sha
					cs, err := api.getCommitStats(repo.Organization, repo.Repository, sha)
					if err != nil {
						log.Errorf(api.ctx, "Error getting stats: %v", err)
						return
					}
					statsChannel <- cs
				}
			}
			wg.Done()
		}()
	}

	// list organizations: GET /user/orgs
	go func() {
		organizations, err := api.getOrganizations()
		if err != nil {
			log.Errorf(api.ctx, "ERROR GETTING ORGANIZATIONS: %v", err)
			return
		}
		for _, organization := range organizations {
			// list repositories: GET /orgs/:org/repos
			repositories, err := api.getRepositories(organization)
			if err != nil {
				log.Errorf(api.ctx, "ERROR GETTING REPOSITORIES: %v", err)
				return
			}
			log.Infof(api.ctx, "REPOSITORIES:%v", repositories)
			for _, repository := range repositories {
				repositoryChannel <- Repo{organization, repository}
			}
		}
		close(repositoryChannel)
	}()

	go func() {
		wg.Wait()
		close(statsChannel)
	}()

	for cs := range statsChannel {
		stats.Additions += cs.Additions
		stats.Deletions += cs.Deletions
	}

	return stats, nil
}

func (api *GithubAPI) getOrganizations() ([]string, error) {
	endpoint := "/user/orgs"
	var data []struct {
		Login string `json:"login"`
	}
	err := api.makeAPIRequest(endpoint, nil, &data)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(data))
	for i, v := range data {
		names[i] = v.Login
	}

	return names, nil
}

func (api *GithubAPI) getRepositories(organization string) ([]string, error) {
	// GET /orgs/:org/repos
	var data []struct {
		Name string `json:"name"`
	}
	err := api.makeAPIRequest("/orgs/"+organization+"/repos", nil, &data)
	if err != nil {
		return nil, err
	}
	names := make([]string, len(data))
	for i, v := range data {
		names[i] = v.Name
	}
	return names, nil
}

func (api *GithubAPI) getUserCommitShas(organization, repository string, since time.Time) ([]string, error) {
	values := make(url.Values)
	values.Add("author", api.username)
	values.Add("since", since.Format(time.RFC3339))
	// GET /repos/:owner/:repo/commits
	endpoint := "/repos/" + organization + "/" + repository + "/commits"
	var data []struct {
		SHA string `json:"sha"`
	}
	err := api.makeAPIRequest(endpoint, values, &data)
	if err != nil {
		return nil, err
	}
	shas := make([]string, len(data))
	for i, v := range data {
		shas[i] = v.SHA
	}
	return shas, nil
}

func (api *GithubAPI) getCommitStats(organization, repository, sha string) (CommitStats, error) {
	var stats CommitStats
	endpoint := "/repos/" + organization + "/" + repository + "/commits/" + sha
	var data struct {
		Stats struct {
			Additions int `json:"additions"`
			Deletions int `json:"deletions"`
		} `json:"stats"`
	}
	err := api.makeAPIRequest(endpoint, nil, &data)
	if err != nil {
		return stats, err
	}
	stats.Additions = data.Stats.Additions
	stats.Deletions = data.Stats.Deletions
	return stats, nil
}

func (api *GithubAPI) makeAPIRequest(endpoint string, values url.Values, dst interface{}) error {
	client := urlfetch.Client(api.ctx)
	if values == nil {
		values = make(url.Values)
	}
	values.Add("access_token", api.accessToken)
	// GET /user/orgs
	response, err := client.Get(githubAPIURL + endpoint + "?" + values.Encode())
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bs, _ := ioutil.ReadAll(response.Body)
	//log.Infof(api.ctx, "GET %s RESPONSE: %s", endpoint, string(bs))
	return json.Unmarshal(bs, dst)
}
