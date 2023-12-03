package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Shemistan/uzum_admin/internal/db"
	"github.com/Shemistan/uzum_admin/internal/models"
)

func (s *AdminService) AddProduct(ctx context.Context, req *models.Product) error {
	user, err := s.GetUser(ctx)
	if err != nil {
		return err
	}
	if user.Role != "admin" {
		return errors.New("access denied")
	}
	return s.storage.CreateProduct(ctx, db.CreateProductParams{
		Name:        req.Name,
		Price:       sql.NullFloat64{Float64: req.Price, Valid: true},
		Description: sql.NullString{String: req.Description, Valid: true},
		Quantity:    req.Quantity,
	})
}
