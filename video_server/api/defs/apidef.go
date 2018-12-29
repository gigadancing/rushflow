package defs

// requests
type UserCredential struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

// Data model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

//
type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}
