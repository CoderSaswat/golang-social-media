package service_impl

import (
	"gorm.io/gorm"
	"social-media/pkg/model"
	"social-media/pkg/repository"
	"social-media/pkg/service"
)

type ProductServiceImpl struct {
	userRepository *repository.UserRepository
}

func (p ProductServiceImpl) GetOrdersFromProduct(productId int) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductServiceImpl) GetAllProducts() []model.Product {
	//TODO implement me
	panic("implement me")
}

func (p ProductServiceImpl) GetProductById() model.Product {
	//TODO implement me
	panic("implement me")
}

func (p ProductServiceImpl) CreateProduct(product model.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductServiceImpl) UpdateProduct(product model.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductServiceImpl) DeleteProduct(productId int) error {
	//TODO implement me
	panic("implement me")
}

func NewProductServiceImpl(db *gorm.DB) service.ProductService {
	return &ProductServiceImpl{userRepository: repository.NewUserRepository(db)}
}
