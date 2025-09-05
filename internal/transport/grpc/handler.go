package grpc

import (
	"context"
	userpb "github.com/TwiLightDM/project-protos/proto/user"
	"github.com/TwiLightDM/users-service/internal/user"
)

type Handler struct {
	svc user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(_ context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	u := user.User{
		Email:    req.Email,
		Password: req.Password,
	}
	created, err := h.svc.CreateUser(u)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:       created.Id,
			Email:    created.Email,
			Password: created.Password,
		},
	}, nil
}

func (h *Handler) GetUser(_ context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	found, err := h.svc.ReadUserById(req.Id)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:       found.Id,
			Email:    found.Email,
			Password: found.Password,
		},
	}, nil
}

func (h *Handler) ListUsers(_ context.Context, _ *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.ReadAllUsers()
	if err != nil {
		return nil, err
	}

	result := make([]*userpb.User, 0, len(users))
	for _, u := range users {
		result = append(result, &userpb.User{
			Id:       u.Id,
			Email:    u.Email,
			Password: u.Password,
		})
	}

	return &userpb.ListUsersResponse{Users: result}, nil
}

func (h *Handler) UpdateUser(_ context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	u := user.User{}
	if req.Email != nil {
		u.Email = *req.Email
	}
	if req.Password != nil {
		u.Password = *req.Password
	}

	updated, err := h.svc.UpdateUser(req.Id, u)
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:       updated.Id,
			Email:    updated.Email,
			Password: updated.Password,
		},
	}, nil
}

func (h *Handler) DeleteUser(_ context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	if err := h.svc.DeleteUser(req.Id); err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResponse{}, nil
}
