package repository

import (
	"gorm.io/gorm"
	"social-media/pkg/model"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) CreatePost(post *model.Post) error {
	return r.DB.Create(post).Error
}

func (r *PostRepository) GetPostByID(postID uint) (*model.Post, error) {
	var post model.Post
	err := r.DB.Preload("Images").Preload("User").Preload("User.Address").First(&post, postID).Error
	return &post, err
}

func (r *PostRepository) GetAllPosts() ([]model.Post, error) {
	var posts []model.Post
	err := r.DB.Preload("Images").Preload("User").Preload("User.Address").Find(&posts).Error
	return posts, err
}

func (r *PostRepository) UpdatePost(post *model.Post) error {
	return r.DB.Save(post).Error
}

func (r *PostRepository) DeletePost(postID uint) error {
	return r.DB.Delete(&model.Post{}, postID).Error
}

func (r *PostRepository) GetAllPostsByUserID(userId uint) ([]model.Post, error) {
	var posts []model.Post

	// Write your raw SQL query
	//query := `
	//    SELECT posts.*, images.*, users.*, addresses.*
	//    FROM posts
	//    LEFT JOIN images ON posts.id = images.post_id
	//    LEFT JOIN users ON posts.user_id = users.id
	//    LEFT JOIN addresses ON users.id = addresses.user_id
	//    WHERE users.id = ?
	//`
	//
	//// Execute the raw SQL query
	//rows, err := r.DB.Raw(query, userId).Rows()
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//
	//// Iterate through the result rows and manually map to Post struct
	//for rows.Next() {
	//	var post model.Post
	//	var image model.Image
	//	var user model.User
	//	var address model.Address
	//
	//	if err := r.DB.ScanRows(rows, &post); err != nil {
	//		return nil, err
	//	}
	//	if err := r.DB.ScanRows(rows, &image); err != nil {
	//		return nil, err
	//	}
	//	if err := r.DB.ScanRows(rows, &user); err != nil {
	//		return nil, err
	//	}
	//	if err := r.DB.ScanRows(rows, &address); err != nil {
	//		return nil, err
	//	}
	//
	//	// Assign nested objects to the post struct
	//	post.Images = append(post.Images, image)
	//	post.User = user
	//	post.User.Address = address
	//
	//	// Append the post to the result slice
	//	posts = append(posts, post)
	//}

	// Write your raw SQL query
	//query := `
	//    SELECT *
	//    FROM posts
	//    LEFT JOIN images ON posts.id = images.post_id
	//    LEFT JOIN users ON posts.user_id = users.id
	//    LEFT JOIN addresses ON users.id = addresses.user_id
	//    WHERE users.id = ?
	//`

	// Execute the raw SQL query
	//err = r.DB.Raw(query, userId).Scan(&posts).Error

	err := r.DB.Preload("Images").Preload("User").Preload("User.Address").Where("user_id = ?", userId).Find(&posts).Error
	return posts, err
}
