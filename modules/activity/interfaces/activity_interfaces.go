package interfaces

import (
	"github.com/fahmiaz411/devcode/modules/activity/domain"

	"context"

	"github.com/gofiber/fiber/v2"
)

type ActivityUsecase interface {
	Create(c *fiber.Ctx, req domain.ActivityCreateRequest) (res domain.ActivityCreateResponse, err error)
	Update(c *fiber.Ctx, req domain.ActivityUpdateRequest) (res domain.ActivityUpdateResponse, err error)
	Delete(c *fiber.Ctx, req domain.ActivityDeleteRequest) (res domain.ActivityDeleteResponse, err error)
	GetAll(c *fiber.Ctx, req domain.ActivityGetAllRequest) (res domain.ActivityGetAllResponse, err error)
	GetOne(c *fiber.Ctx, req domain.ActivityGetOneRequest) (res domain.ActivityGetOneResponse, err error)
}

type ActivityRepoMysql interface {
	Create(ctx context.Context, req domain.ActivityCreateRequest) (res domain.ActivityCreateResponse, err error)
	Update(ctx context.Context, req domain.ActivityUpdateRequest) (res domain.ActivityUpdateResponse, err error)
	Delete(ctx context.Context, req domain.ActivityDeleteRequest) (res domain.ActivityDeleteResponse, err error)
	GetAll(ctx context.Context, req domain.ActivityGetAllRequest) (res domain.ActivityGetAllResponse, err error)
	GetOne(ctx context.Context, req domain.ActivityGetOneRequest) (res domain.ActivityGetOneResponse, err error)
}