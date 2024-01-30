package service

import "social-media/pkg/model"

type ProductService interface {
	GetAllProducts() []model.Product
	GetProductById() model.Product
	CreateProduct(product model.Product) error
	UpdateProduct(product model.Product) error
	DeleteProduct(productId int) error
	GetOrdersFromProduct(productId int) error
}
