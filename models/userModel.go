package models

//* User Model representation
//? Initially the User Model is a JSON until the DB is implemented.
//? DB plan is MongoDB which will change to BSON probably

type User struct {
	Username string `bson:"username"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
