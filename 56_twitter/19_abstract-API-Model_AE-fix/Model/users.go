package Model

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string
}
