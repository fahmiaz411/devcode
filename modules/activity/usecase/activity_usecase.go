package usecase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/fahmiaz411/devcode/helper/constant"
	"github.com/fahmiaz411/devcode/helper/message"
	"github.com/fahmiaz411/devcode/helper/web"
	"github.com/fahmiaz411/devcode/modules/activity/domain"
	"github.com/fahmiaz411/devcode/modules/activity/interfaces"
	"github.com/fahmiaz411/devcode/modules/activity/repository"

	"github.com/gofiber/fiber/v2"
)

type Usecase struct {
	repo           *repository.Repository
	contentTimeout time.Duration
}

func NewUsecase(repo *repository.Repository, timeout time.Duration) interfaces.ActivityUsecase {
	return &Usecase{
		repo:           repo,
		contentTimeout: timeout,
	}
}


func (u *Usecase) Create(c *fiber.Ctx, req domain.ActivityCreateRequest) (res domain.ActivityCreateResponse, err error) {
	ctx, cancel := context.WithTimeout(c.UserContext(), u.contentTimeout)
	defer cancel()

	res, err = u.repo.MySQL.Create(ctx, req)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
			Data: struct{}{},
		})
		return
	}

	return
}

func (u *Usecase) Update(c *fiber.Ctx, req domain.ActivityUpdateRequest) (res domain.ActivityUpdateResponse, err error) {
	ctx, cancel := context.WithTimeout(c.UserContext(), u.contentTimeout)
	defer cancel()

	var activity domain.ActivityGetOneResponse
	activity, err = u.GetOne(c, domain.ActivityGetOneRequest{
		ID: req.ID,
	})
	if err != nil {
		return
	}

	req.UpdatedAt = time.Now().UTC()

	res, err = u.repo.MySQL.Update(ctx, req)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
			Data: struct{}{},
		})
		return
	}

	res.ID = activity.ID
	res.Title = req.Title
	res.Email = activity.Email
	res.CreatedAt = activity.CreatedAt
	res.UpdatedAt = req.UpdatedAt
	res.DeletedAt = activity.DeletedAt

	return 
}

func (u *Usecase) Delete(c *fiber.Ctx, req domain.ActivityDeleteRequest) (res domain.ActivityDeleteResponse, err error) {
	ctx, cancel := context.WithTimeout(c.UserContext(), u.contentTimeout)
	defer cancel()

	_, err = u.GetOne(c, domain.ActivityGetOneRequest{
		ID: req.ID,
	})
	if err != nil {
		return
	}
	
	res, err = u.repo.MySQL.Delete(ctx, req)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
			Data: struct{}{},
		})
	}

	return 
}

func (u *Usecase) GetAll(c *fiber.Ctx, req domain.ActivityGetAllRequest) (res domain.ActivityGetAllResponse, err error) {
	ctx, cancel := context.WithTimeout(c.UserContext(), u.contentTimeout)
	defer cancel()

	res, err = u.repo.MySQL.GetAll(ctx, req)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
			Data: struct{}{},
		})
		return
	}

	return
}

func (u *Usecase) GetOne(c *fiber.Ctx, req domain.ActivityGetOneRequest) (res domain.ActivityGetOneResponse, err error) {
	ctx, cancel := context.WithTimeout(c.UserContext(), u.contentTimeout)
	defer cancel()

	res, err = u.repo.MySQL.GetOne(ctx, req)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
			Data: struct{}{},
		})
		return
	} else if res.ID == int64(constant.ZeroValue) {
		c.Status(http.StatusNotFound).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusNotFound),
			Message: message.NotFound(domain.Model, "ID", fmt.Sprint(req.ID)),
			Data: struct{}{},
		})
		err = fmt.Errorf(http.StatusText(http.StatusNotFound))
		return
	}

	return 
}