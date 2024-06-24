package model

type Comment struct {
	ID      int        `json:"id"`
	Author  string     `json:"author"`
	Content string     `json:"content"`
	Post    int        `json:"post" db:"postid"`
	Replies []*Comment `json:"replies,omitempty"`
	ReplyTo *int       `json:"replyTo,omitempty"`
}

type CreateCommentReq struct {
	Author  string `json:"author"`
	Content string `json:"content"`
	Post    int    `json:"post"`
	ReplyTo *int   `json:"replyTo,omitempty"`
}

type CreatePostReq struct {
	Name            string `json:"name"`
	Content         string `json:"content"`
	Author          string `json:"author"`
	CommentsAllowed bool   `json:"commentsAllowed"`
}

type Post struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Author          string     `json:"author"`
	Content         string     `json:"content"`
	CommentsAllowed bool       `json:"commentsAllowed"`
	Comments        []*Comment `json:"comments,omitempty"`
}

type PostListEl struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
