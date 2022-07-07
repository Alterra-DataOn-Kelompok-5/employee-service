package employee

import (
	"context"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/repository"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/constant"
	pkgdto "github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/dto"
	res "github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/util/response"
)

type service struct {
	EmployeeRepository repository.Employee
}

type Service interface {
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.EmployeeResponse], error)
	FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.EmployeeDetailResponse, error)
	UpdateById(ctx context.Context, payload *dto.UpdateEmployeeRequestBody) (*dto.EmployeeDetailResponse, error)
	DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.EmployeeWithCUDResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		EmployeeRepository: f.EmployeeRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.EmployeeResponse], error) {

	employees, info, err := s.EmployeeRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data []dto.EmployeeResponse

	for _, employee := range employees {
		data = append(data, dto.EmployeeResponse{
			ID:       employee.ID,
			Fullname: employee.Fullname,
			Email:    employee.Email,
		})

	}

	result := new(pkgdto.SearchGetResponse[dto.EmployeeResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.EmployeeDetailResponse, error) {
	data, err := s.EmployeeRepository.FindByID(ctx, payload.ID, true)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.EmployeeDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.EmployeeDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.EmployeeDetailResponse{
		EmployeeResponse: dto.EmployeeResponse{
			ID:       data.ID,
			Fullname: data.Fullname,
			Email:    data.Email,
		},
		Role: dto.RoleResponse{
			ID:   data.Role.ID,
			Name: data.Role.RoleName,
		},
		Division: dto.DivisionResponse{
			ID:   data.Division.ID,
			Name: data.Division.DivisionName,
		},
	}

	return result, nil
}

func (s *service) UpdateById(ctx context.Context, payload *dto.UpdateEmployeeRequestBody) (*dto.EmployeeDetailResponse, error) {
	employee, err := s.EmployeeRepository.FindByID(ctx, *payload.ID, false)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.EmployeeDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.EmployeeDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	_, err = s.EmployeeRepository.Edit(ctx, &employee, payload)
	if err != nil {
		return &dto.EmployeeDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.EmployeeDetailResponse{
		EmployeeResponse: dto.EmployeeResponse{
			ID:       employee.ID,
			Fullname: employee.Fullname,
			Email:    employee.Email,
		},
		Role: dto.RoleResponse{
			ID:   employee.Role.ID,
			Name: employee.Role.RoleName,
		},
		Division: dto.DivisionResponse{
			ID:   employee.Division.ID,
			Name: employee.Division.DivisionName,
		},
	}

	return result, nil
}

func (s *service) DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.EmployeeWithCUDResponse, error) {
	employee, err := s.EmployeeRepository.FindByID(ctx, payload.ID, false)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.EmployeeWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.EmployeeWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	_, err = s.EmployeeRepository.Destroy(ctx, &employee)
	if err != nil {
		return &dto.EmployeeWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.EmployeeWithCUDResponse{
		EmployeeResponse: dto.EmployeeResponse{
			ID:       employee.ID,
			Fullname: employee.Fullname,
			Email:    employee.Email,
		},
		CreatedAt: employee.CreatedAt,
		UpdatedAt: employee.UpdatedAt,
		DeletedAt: employee.DeletedAt,
	}

	return result, nil
}
