package search

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/search"
)

func init() {
	http.HandleFunc("/put", handlePut)
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/search", handleSearch)
}

type Document struct {
	Name string
	Text string
}

func handleSearch(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	index, err := search.Open("example")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	iterator := index.Search(ctx, "the", nil)
	for {
		var document Document
		id, err := iterator.Next(&document)
		if err == search.Done {
			break
		} else if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		fmt.Fprintf(res, "ID: %v, DOCUMENT: %v \n\n", id, document)
	}
}

func handleGet(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	index, err := search.Open("example")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	var document Document
	err = index.Get(ctx, "23687ac5-5417-4b13-b9f7-4e674fa5c227", &document)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	fmt.Fprintf(res, "%v", document.Text)
}

func handlePut(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	index, err := search.Open("example")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	var document = &Document{
		Name: "example",
		Text: `It is a truth universally acknowledged, that a single man in possession of a good fortune, must be in want of a wife.
However little known the feelings or views of such a man may be on his first entering a neighbourhood, this truth is so well fixed in the minds of the surrounding families, that he is considered the rightful property of some one or other of their daughters.
"My dear Mr. Bennet," said his lady to him one day, "have you heard that Netherfield Park is let at last?"
Mr. Bennet replied that he had not.
"But it is," returned she; "for Mrs. Long has just been here, and she told me all about it."
Mr. Bennet made no answer.
"Do you not want to know who has taken it?" cried his wife impatiently.
"You want to tell me, and I have no objection to hearing it."
This was invitation enough.
"Why, my dear, you must know, Mrs. Long says that Netherfield is taken by a young man of large fortune from the north of England; that he came down on Monday in a chaise and four to see the place, and was so much delighted with it, that he agreed with Mr. Morris immediately; that he is to take possession before Michaelmas, and some of his servants are to be in the house by the end of next week."
"What is his name?"
"Bingley."
"Is he married or single?"
"Oh! Single, my dear, to be sure! A single man of large fortune; four or five thousand a year. What a fine thing for our girls!"
"How so? How can it affect them?"
"My dear Mr. Bennet," replied his wife, "how can you be so tiresome! You must know that I am thinking of his marrying one of them."
"Is that his design in settling here?"
"Design! Nonsense, how can you talk so! But it is very likely that he may fall in love with one of them, and therefore you must visit him as soon as he comes."
"I see no occasion for that. You and the girls may go, or you may send them by themselves, which perhaps will be still better, for as you are as handsome as any of them, Mr. Bingley may like you the best of the party."
"My dear, you flatter me. I certainly have had my share of beauty, but I do not pretend to be anything extraordinary now. When a woman has five grown-up daughters, she ought to give over thinking of her own beauty."
"In such cases, a woman has not often much beauty to think of."
"But, my dear, you must indeed go and see Mr. Bingley when he comes into the neighbourhood."
"It is more than I engage for, I assure you."
"But consider your daughters. Only think what an establishment it would be for one of them. Sir William and Lady Lucas are determined to go, merely on that account, for in general, you know, they visit no newcomers. Indeed you must go, for it will be impossible for us to visit him if you do not."
"You are over-scrupulous, surely. I dare say Mr. Bingley will be very glad to see you; and I will send a few lines by you to assure him of my hearty consent to his marrying whichever he chooses of the girls; though I must throw in a good word for my little Lizzy."
"I desire you will do no such thing. Lizzy is not a bit better than the others; and I am sure she is not half so handsome as Jane, nor half so good-humoured as Lydia. But you are always giving her the preference."`,
	}

	//	var document = &Document{
	//		Name: "StephenKing",
	//		Text: `Jacobs's electrical workshop was in West Tulsa. I don't know what that part of town is like now, but in 1992 it was a forlorn industrial zone where a lot of the industries seemed to be dead or dying. He pulled into the parking lot of an all-but-destitute strip mall on Olympia Avenue and parked in front of Wilson Auto Body. "It was standing empty for a long time, that's what the Realtor told me," Jacobs said. He was dressed in faded jeans and a blue golf shirt, his hair washed and combed, his eyes sparkling with excitement. Just looking at him made me nervous. "I had to take a year's lease, but it was still dirt cheap. Come on in." "You ought to take down the sign and put up your own," I said. I framed it with hands that were only shaking a little. "'Portraits in Lightning, C. D. Jacobs, Proprietor.' It would look good." "I won't be in Tulsa that long," he said, "and the portraits are really just a way of supporting myself while I conduct my experiments. I've come a long way since my pastoral days, but I've still got a long way to go. You have no idea. Come in, Jamie. Come in." He unlocked a door and led me through an office that was empty of furniture, although I could still see square clean patches on the grimy linoleum, where the legs of a desk had once stood. On the wall was a curling calendar with April 1989 showing. The garage had a corrugated metal roof and I expected it to be baking under the September sun, but it was wonderfully cool. I could hear the whisper of air conditioners. When he flicked a bank of switches—recently modified, judging from the makeshift way the wires stuck out of the uncovered holes where the plates had been—a dozen brilliant lights came on. If not for the oil-darkened concrete and the rectangular caverns where two lifts had once been, you would have thought it was an operating theater. "It must cost a fortune to air-condition this place," I said. "Especially when you've got all those lights blazing." "Dirt cheap. The air conditioners are my own design. They draw very little power, and most of that I generate myself. I could gener- ate all of it, but I wouldn't want Tulsa Power and Light down here, snooping around to fifind out if I was volt-jacking, somehow. As for the lights . . . you could wrap a hand around one of the bulbs with- out burning yourself. Or even heating your skin, for that matter." Our footfalls echoed in all that empty space. So did our voices. It was like being in the company of phantoms. It just feels that way because I'm strung out, I told myself. "Listen, Charlie — you're not messing with anything radioactive, are you?" He grimaced and shook his head. "Nuclear's the last thing I'm interested in. It's energy for idiots. A dead end." "So how do you generate the juice?" "Electricity breeds electricity, if you know what you're doing. Leave it at that. Step over here, Jamie." There were three or four long tables at the end of the room with electrical stuff on them. I recognized an oscilloscope, a spectrom- eter, and a couple of things that looked like Marshall amps but could have been batteries of some kind. There was a control board that looked mostly torn apart, and several stacked consoles with darkened dials. Thick electrical cords snaked every whichway. Some disappeared into closed metal containers that could have been Craftsman tool chests; others just looped back to the dark equip- ment. This could all be a fantasy, I thought. Equipment that only comes alive in his imagination. But the Portraits in Lightning weren't make- believe. I had no idea how he was making those, his explanation had been vague at best, but he was making them. And although I was standing directly beneath one of those brilliant lights, it really did not seem to be throwing any heat. "There doesn't seem to be much here," I said doubtfully. "I expected more." "Flashing lights! Chrome-plated knife-switches sticking out of science fiction control panels! Star Trek telescreens! Possibly a tele- portation chamber, or a hologram of Noah's Ark in a cloud cham- ber!" He laughed cheerily. "Nothing like that," I said, although he had pretty much hit the nail on the head. "It just seems kind of . . . sparse." "It is. I've gone about as far as I can for the time being. I've sold some of my equipment. Other stuff—more controversial stuff— I've dismantled and put in storage. I've done good work in Tulsa, especially considering how little spare time I have. Keeping body and soul together is an annoying business, as I suppose you know." I certainly did. "But yes, I made some progress toward my ultimate goal. Now I need to think, and I don't believe I can do that when I'm turning half a dozen tips a night." "Your ultimate goal being what?" He ignored the question this time, too. "Step over here, Jamie. Would you like a small pick-me-up before we begin?" I wasn't sure I wanted to begin, but I wanted a pick-me-up, all right. Not for the first time, I considered just snatching the little brown bottle and running. Only he'd probably catch me and wrest it away. I was younger, and almost over the flu, but he was still in better shape. He hadn't suffered a shattered hip and leg in a motor- cycle accident, for one thing. He grabbed a paint-spattered wooden chair and set it in front of one of the black boxes that looked like a Marshall amp. "Sit here." But I didn't, not right away. There was a picture on one of the tables, the kind with a little wedge on the back to prop it up. He saw me reach for it and made a move as if to stop me. Then he just stood there. A song on the radio can bring back the past with fierce (if mercifully transitory) immediacy: a first kiss, a good time with your buddies, or an unhappy life-passage. I can never hear Fleetwood Mac's "Go Your Own Way" without thinking of my mother's last painful weeks; that spring it seemed to be on the radio every time I turned it on. A picture can have the same effect. I looked at this one and all at once I was eight again. My sister was helping Morrie set up dominos in Toy Corner while Patsy Jacobs played "Bringing in the Sheaves," swaying on the piano bench, her smooth blond hair shifting from side to side. It was a studio portrait. Patsy was wearing the sort of billowy, shin-length dress that went out of fashion years ago, but it looked good on her. The kid was on her lap, wearing short pants and a sweater vest. A cowlick I remembered well stuck up at the back of his head. Read more: http://www.rollingstone.com/culture/features/stephen-king-exclusive-read-an-excerpt-from-new-book-revival-20141027#ixzz3rGP0EkMZ Follow us: @rollingstone on Twitter | RollingStone on Facebook`,
	//	}

	id, err := index.Put(ctx, "", document)

	fmt.Fprintf(res, "ID: %v, ERROR: %v", id, err)
}
