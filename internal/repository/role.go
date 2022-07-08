package repository

import (
	"context"
	"strings"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/model"
	pkgdto "github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/dto"
	"gorm.io/gorm"
)

type Role interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, p *pkgdto.Pagination) ([]model.Role, *pkgdto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.Role, error)
	Save(ctx context.Context, role *dto.CreateRoleRequestBody) (model.Role, error)
	Edit(ctx context.Context, oldrole *model.Role, updateData *dto.UpdateRoleRequestBody) (*model.Role, error)
	Destroy(ctx context.Context, role *model.Role) (*model.Role, error)
	ExistByName(ctx context.Context, name string) (bool, error)
}

type role struct {
	Db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *role {
	return &role{
		db,
	}
}

func (r *role) FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Role, *pkgdto.PaginationInfo, error) {
	var roles []model.Role
	var count int64

	query := r.Db.WithContext(ctx).Model(&model.Role{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?", search, search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := pkgdto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&roles).Error

	return roles, pkgdto.CheckInfoPagination(pagination, count), err
}

func (r *role) FindByID(ctx context.Context, id uint) (model.Role, error) {
	var role model.Role
	if err := r.Db.WithContext(ctx).Model(&model.Role{}).Where("id = ?", id).First(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}

func (r *role) Save(ctx context.Context, role *dto.CreateRoleRequestBody) (model.Role, error) {
	newRole := model.Role{
		Name: *role.Name,
	}
	if err := r.Db.WithContext(ctx).Save(&newRole).Error; err != nil {
		return newRole, err
	}
	return newRole, nil
}

func (r *role) Edit(ctx context.Context, oldRole *model.Role, updateData *dto.UpdateRoleRequestBody) (*model.Role, error) {
	if updateData.Name != nil {
		oldRole.Name = *updateData.Name
	}

	if err := r.Db.WithContext(ctx).Save(oldRole).Find(oldRole).Error; err != nil {
		return nil, err
	}

	return oldRole, nil
}

func (r *role) Destroy(ctx context.Context, role *model.Role) (*model.Role, error) {
	if err := r.Db.WithContext(ctx).Delete(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *role) ExistByName(ctx context.Context, name string) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.Role{}).Where("name = ?", name).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}
