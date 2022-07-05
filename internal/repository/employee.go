package repository

import (
	"context"
	"strings"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/model"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/util"
	"gorm.io/gorm"
)

type Employee interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, p *dto.Pagination) ([]model.Employee, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.Employee, error)
	FindByEmail(ctx context.Context, email *string) (*model.Employee, error)
	ExistByEmail(ctx context.Context, email *string) (bool, error)
	ExistByID(ctx context.Context, id uint) (bool, error)
	Save(ctx context.Context, employee *dto.RegisterEmployeeRequestBody) (model.Employee, error)
	Edit(ctx context.Context, oldEmployee model.Employee, updateData *dto.UpdateEmployeeRequestBody) (*model.Employee, error)
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

func (r *employee) ExistByEmail(ctx context.Context, email *string) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.Employee{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}

func (r *employee) ExistByID(ctx context.Context, id uint) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.Employee{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}

func (r *employee) Save(ctx context.Context, employee *dto.RegisterEmployeeRequestBody) (model.Employee, error) {
	newEmployee := model.Employee{
		Fullname:   employee.Fullname,
		Email:      employee.Email,
		Password:   employee.Password,
		RoleID:     *employee.RoleID,
		DivisionID: *employee.DivisionID,
	}
	if err := r.Db.WithContext(ctx).Save(&newEmployee).Error; err != nil {
		return newEmployee, err
	}
	return newEmployee, nil
}

func (r *employee) Edit(ctx context.Context, oldEmployee model.Employee, updateData *dto.UpdateEmployeeRequestBody) (*model.Employee, error) {
	if updateData.Fullname != nil {
		oldEmployee.Fullname = *updateData.Fullname
	}
	if updateData.Email != nil {
		oldEmployee.Email = *updateData.Email
	}
	if updateData.Password != nil {
		hashedPassword, err := util.HashPassword(*updateData.Password)
		if err != nil {
			return &oldEmployee, err
		}
		oldEmployee.Password = hashedPassword
	}
	if updateData.DivisionID != nil {
		oldEmployee.DivisionID = *updateData.DivisionID
	}
	if updateData.RoleID != nil {
		oldEmployee.RoleID = *updateData.RoleID
	}

	if err := r.Db.WithContext(ctx).Save(&oldEmployee).Error; err != nil {
		return &oldEmployee, err
	}

	var employee model.Employee
	if err := r.Db.WithContext(ctx).Preload("Division").Preload("Role").Where("id = ?", oldEmployee.ID).Find(&employee).Error; err != nil {
		return &oldEmployee, err
	}

	return &employee, nil
}
