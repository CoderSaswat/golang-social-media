package model

import (
	"gorm.io/gorm"
)

type Address struct {
	Village string `json:"village"`
	Pin     string `json:"pin"`
	State   string `json:"state"`
	Country string `json:"country"`
	UserId  uint   `json:"userId"`
}

type User struct {
	gorm.Model
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
	Address Address `json:"address"`
}

type Post struct {
	gorm.Model
	Caption string  `json:"caption"`
	Images  []Image `json:"images"`
	UserID  uint    `json:"userId"`
	User    User    `json:"user"`
}

type Image struct {
	gorm.Model
	URL    string `json:"url"`
	PostID uint   `json:"postID"`
}

type Comment struct {
	gorm.Model
	Text   string `json:"text"`
	PostID uint   `json:"postID"`
	Post   Post   `json:"post"`
	UserID uint   `json:"userID"`
	User   User   `json:"user"`
}

type Like struct {
	gorm.Model
	UserID uint `json:"userID"`
	User   User `json:"user"`
	PostID uint `json:"postID"`
	Post   Post `json:"post"`
}

type Notification struct {
	gorm.Model
	Message  string `json:"message"`
	UserID   uint   `json:"userID"`
	User     User   `json:"user"`
	SenderID uint   `json:"senderID"`
	Sender   User   `json:"sender"`
	PostID   uint   `json:"postID"`
	Post     Post   `json:"post"`
}
type Follow struct {
	gorm.Model
	UserID       uint   `json:"userID"`
	User         User   `json:"user"`
	TargetUserID uint   `json:"targetUserID"`
	TargetUser   User   `json:"targetUser"`
	FollowType   string `json:"followType"`
}
