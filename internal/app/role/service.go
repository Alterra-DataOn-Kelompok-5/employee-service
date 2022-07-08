package role

import (
	"context"
	"errors"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/repository"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/constant"
	pkgdto "github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/dto"
	res "github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/util/response"
)

type service struct {
	RoleRepository repository.Role
}

type Service interface {
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.RoleResponse], error)
	FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoleResponse, error)
	Store(ctx context.Context, payload *dto.CreateRoleRequestBody) (*dto.RoleResponse, error)
	UpdateById(ctx context.Context, payload *dto.UpdateRoleRequestBody) (*dto.RoleResponse, error)
	DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoleWithCUDResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		RoleRepository: f.RoleRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.RoleResponse], error) {
	roles, info, err := s.RoleRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data []dto.RoleResponse

	for _, role := range roles {
		data = append(data, dto.RoleResponse{
			ID:   role.ID,
			Name: role.Name,
		})

	}

	result := new(pkgdto.SearchGetResponse[dto.RoleResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}
func (s *service) FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoleResponse, error) {
	var result dto.RoleResponse
	data, err := s.RoleRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.RoleResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.RoleResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = data.ID
	result.Name = data.Name

	return &result, nil
}

func (s *service) Store(ctx context.Context, payload *dto.CreateRoleRequestBody) (*dto.RoleResponse, error) {
	var result dto.RoleResponse
	isExist, err := s.RoleRepository.ExistByName(ctx, *payload.Name)
	if err != nil {
		return &result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if isExist {
		return &result, res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("role already exists"))
	}

	data, err := s.RoleRepository.Save(ctx, payload)
	if err != nil {
		return &result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = data.ID
	result.Name = data.Name

	return &result, nil
}

func (s *service) UpdateById(ctx context.Context, payload *dto.UpdateRoleRequestBody) (*dto.RoleResponse, error) {
	role, err := s.RoleRepository.FindByID(ctx, *payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.RoleResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.RoleResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	_, err = s.RoleRepository.Edit(ctx, &role, payload)
	if err != nil {
		return &dto.RoleResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	var result dto.RoleResponse
	result.ID = role.ID
	result.Name = role.Name

	return &result, nil
}
func (s *service) DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.RoleWithCUDResponse, error) {
	role, err := s.RoleRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.RoleWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.RoleWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	_, err = s.RoleRepository.Destroy(ctx, &role)
	if err != nil {
		return &dto.RoleWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.RoleWithCUDResponse{
		RoleResponse: dto.RoleResponse{
			ID:   role.ID,
			Name: role.Name,
		},
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
		DeletedAt: role.DeletedAt,
	}

	return result, nil
}
