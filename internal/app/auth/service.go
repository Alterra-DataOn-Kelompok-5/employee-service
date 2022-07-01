package auth

import (
	"context"
	"errors"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/repository"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/constant"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/util"
	res "github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/util/response"
)

type service struct {
	EmployeeRepository repository.Employee
}

type Service interface {
	LoginByEmailAndPassword(ctx context.Context, payload *dto.ByEmailAndPasswordRequest) (*dto.EmployeeWithJWTResponse, error)
	RegisterByEmailAndPassword(ctx context.Context, payload *dto.RegisterEmployeeRequestBody) (*dto.EmployeeWithJWTResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		EmployeeRepository: f.EmployeeRepository,
	}
}

func (s *service) LoginByEmailAndPassword(ctx context.Context, payload *dto.ByEmailAndPasswordRequest) (*dto.EmployeeWithJWTResponse, error) {
	var result *dto.EmployeeWithJWTResponse

	data, err := s.EmployeeRepository.FindByEmail(ctx, &payload.Email)
	if err != nil {
		if err == constant.RecordNotFound {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	if !(util.CompareHashPassword(payload.Password, data.Password)) {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.EmailOrPasswordIncorrect,
			errors.New(res.ErrorConstant.EmailOrPasswordIncorrect.Response.Meta.Message),
		)
	}
	// TODO: Generate JWT
	result = &dto.EmployeeWithJWTResponse{
		EmployeeResponse: dto.EmployeeResponse{
			ID:       data.ID,
			Fullname: data.Fullname,
			Email:    data.Email,
		},
		JWT: "put jwt token here",
	}

	return result, nil
}

func (s *service) RegisterByEmailAndPassword(ctx context.Context, payload *dto.RegisterEmployeeRequestBody) (*dto.EmployeeWithJWTResponse, error) {
	var result *dto.EmployeeWithJWTResponse
	isExist, err := s.EmployeeRepository.ExistByEmail(ctx, &payload.Email)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if isExist {
		return result, res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("employee already exists"))
	}

	hashedPassword, err := util.HashPassword(payload.Password)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	payload.Password = hashedPassword

	data, err := s.EmployeeRepository.Save(ctx, payload)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// TODO: Generate JWT

	result = &dto.EmployeeWithJWTResponse{
		EmployeeResponse: dto.EmployeeResponse{
			ID:       data.ID,
			Fullname: data.Fullname,
			Email:    data.Email,
		},
		JWT: "put jwt token here",
	}

	return result, nil
}
