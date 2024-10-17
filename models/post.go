package models

// * Structure for Post
type Post struct {
	Title   string `bson:"title"`
	Content string `bson:"content"`
	Author  string `bson:"username"`
}
