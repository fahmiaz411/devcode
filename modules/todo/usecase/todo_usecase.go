package usecase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/fahmiaz411/devcode/helper/constant"
	"github.com/fahmiaz411/devcode/helper/message"
	"github.com/fahmiaz411/devcode/helper/web"
	"github.com/fahmiaz411/devcode/modules/todo/domain"
	"github.com/fahmiaz411/devcode/modules/todo/interfaces"
	"github.com/fahmiaz411/devcode/modules/todo/repository"

	"github.com/gofiber/fiber/v2"
)

type Usecase struct {
	repo           *repository.Repository
	contentTimeout time.Duration
}

func NewUsecase(repo *repository.Repository, timeout time.Duration) interfaces.TodoUsecase {
	return &Usecase{
		repo:           repo,
		contentTimeout: timeout,
	}
}


func (u *Usecase) Create(c *fiber.Ctx, req domain.TodoCreateRequest) (res domain.TodoCreateResponse, err error) {
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

func (u *Usecase) Update(c *fiber.Ctx, req domain.TodoUpdateRequest) (res domain.TodoUpdateResponse, err error) {
	ctx, cancel := context.WithTimeout(c.UserContext(), u.contentTimeout)
	defer cancel()

	var todo domain.TodoGetOneResponse
	todo, err = u.GetOne(c, domain.TodoGetOneRequest{
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

	res.ID = todo.ID
	res.ActivityGroupID = todo.ActivityGroupID
	res.CreatedAt = todo.CreatedAt
	res.UpdatedAt = req.UpdatedAt

	// Title
	if req.Title != constant.EmptyString {
		res.Title = req.Title
	} else {
		res.Title = todo.Title
	}

	// Is Active
	if req.IsActive != nil {
		res.IsActive = *req.IsActive
	} else {
		res.IsActive = todo.IsActive
	}

	// Priority
	if req.Priority != constant.EmptyString {
		res.Priority = req.Priority
	} else {
		res.Priority = todo.Priority
	}

	return 
}

func (u *Usecase) Delete(c *fiber.Ctx, req domain.TodoDeleteRequest) (res domain.TodoDeleteResponse, err error) {
	ctx, cancel := context.WithTimeout(c.UserContext(), u.contentTimeout)
	defer cancel()

	_, err = u.GetOne(c, domain.TodoGetOneRequest{
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

func (u *Usecase) GetAll(c *fiber.Ctx, req domain.TodoGetAllRequest) (res domain.TodoGetAllResponse, err error) {
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

func (u *Usecase) GetOne(c *fiber.Ctx, req domain.TodoGetOneRequest) (res domain.TodoGetOneResponse, err error) {
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
			Message: message.NotFound("Todo", "ID", fmt.Sprint(req.ID)),
			Data: struct{}{},
		})
		err = fmt.Errorf(http.StatusText(http.StatusNotFound))
		return
	}

	return 
}