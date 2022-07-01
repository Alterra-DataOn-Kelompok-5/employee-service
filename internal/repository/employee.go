package repository

import (
	"context"
	"strings"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/model"
	"gorm.io/gorm"
)

type Employee interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, p *dto.Pagination) ([]model.Employee, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.Employee, error)
	FindByEmail(ctx context.Context, email *string) (*model.Employee, error)
	// Save(ctx context.Context, employee *dto.RegisterUserRequestBody) (model.Employee, error)
}

type employee struct {
	Db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *employee {
	return &employee{
		db,
	}
}

func (r *employee) FindAll(ctx context.Context, payload *dto.SearchGetRequest, pagination *dto.Pagination) ([]model.Employee, *dto.PaginationInfo, error) {
	var users []model.Employee
	var count int64

	query := r.Db.WithContext(ctx).Model(&model.Employee{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(fullname) LIKE ? or lower(email) Like ? ", search, search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&users).Error

	return users, dto.CheckInfoPagination(pagination, count), err
}

func (r *employee) FindByID(ctx context.Context, id uint) (model.Employee, error) {
	var user model.Employee
	err := r.Db.WithContext(ctx).Model(&model.Employee{}).Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *employee) FindByEmail(ctx context.Context, email *string) (*model.Employee, error) {
	var data model.Employee
	err := r.Db.WithContext(ctx).Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
