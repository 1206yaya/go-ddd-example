package dtos

type CreateProductRequest struct {
	Name  string `json:"name" validate:"required,min=1,max=100"`
	Price int    `json:"price" validate:"required,gt=0"`
}

type CreateProductResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
