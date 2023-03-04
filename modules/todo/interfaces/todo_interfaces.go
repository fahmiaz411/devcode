package interfaces

import (
	"github.com/fahmiaz411/devcode/modules/todo/domain"

	"context"

	"github.com/gofiber/fiber/v2"
)

type TodoUsecase interface {
	Create(c *fiber.Ctx, req domain.TodoCreateRequest) (res domain.TodoCreateResponse, err error)
	Update(c *fiber.Ctx, req domain.TodoUpdateRequest) (res domain.TodoUpdateResponse, err error)
	Delete(c *fiber.Ctx, req domain.TodoDeleteRequest) (res domain.TodoDeleteResponse, err error)
	GetAll(c *fiber.Ctx, req domain.TodoGetAllRequest) (res domain.TodoGetAllResponse, err error)
	GetOne(c *fiber.Ctx, req domain.TodoGetOneRequest) (res domain.TodoGetOneResponse, err error)
}

type TodoRepoMysql interface {
	Create(ctx context.Context, req domain.TodoCreateRequest) (res domain.TodoCreateResponse, err error)
	Update(ctx context.Context, req domain.TodoUpdateRequest) (res domain.TodoUpdateResponse, err error)
	Delete(ctx context.Context, req domain.TodoDeleteRequest) (res domain.TodoDeleteResponse, err error)
	GetAll(ctx context.Context, req domain.TodoGetAllRequest) (res domain.TodoGetAllResponse, err error)
	GetOne(ctx context.Context, req domain.TodoGetOneRequest) (res domain.TodoGetOneResponse, err error)
}