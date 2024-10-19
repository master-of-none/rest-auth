package models

// * Structure for Post
type Post struct {
	ID      int    `bson:"id"`
	Title   string `bson:"title"`
	Content string `bson:"content"`
	Author  string `bson:"author"`
}

// * Structure for Autoincrement Value in Mongo Database
type PostResult struct {
	SequenceValue int `bson:"sequence_value"`
}
