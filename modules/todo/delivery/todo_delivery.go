package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fahmiaz411/devcode/helper/constant"
	"github.com/fahmiaz411/devcode/helper/field"
	"github.com/fahmiaz411/devcode/helper/message"
	"github.com/fahmiaz411/devcode/helper/params"
	"github.com/fahmiaz411/devcode/helper/query"
	"github.com/fahmiaz411/devcode/helper/slice"
	"github.com/fahmiaz411/devcode/helper/web"
	"github.com/fahmiaz411/devcode/modules/todo/domain"
	"github.com/fahmiaz411/devcode/modules/todo/interfaces"

	"github.com/gofiber/fiber/v2"
)

type RESTHandler struct {
	Usecase interfaces.TodoUsecase
}

func NewRESTHandler(f fiber.Router, usecase interfaces.TodoUsecase) {
	handler := &RESTHandler{
		Usecase: usecase,
	}

	f.Post("/todo-items", handler.Create)

	f.Patch(fmt.Sprintf("/todo-items/:%s", params.TodoId), handler.Update)

	f.Delete(fmt.Sprintf("/todo-items/:%s", params.TodoId), handler.Delete)

	f.Get("/todo-items", handler.GetAll)

	f.Get(fmt.Sprintf("/todo-items/:%s", params.TodoId), handler.GetOne)
}

func (h *RESTHandler) Create(c *fiber.Ctx) error {
	req := domain.TodoCreateRequest{}
	c.BodyParser(&req)

	if req.Title == constant.EmptyString {
		return c.Status(http.StatusBadRequest).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusBadRequest),
			Message: message.CanotNull(field.Title),
			Data: struct{}{},
		})
	} else if req.ActivityGroupID == int64(constant.ZeroValue) {
		return c.Status(http.StatusBadRequest).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusBadRequest),
			Message: message.CanotNull(field.ActivityGroupID),
			Data: struct{}{},
		})
	}

	res, err := h.Usecase.Create(c, req)
	if err != nil {
		return nil
	}

	return c.JSON(web.BaseResponse{
		Status: message.Success,
		Message: message.Success,
		Data: res,
	})
}

func (h *RESTHandler) Update(c *fiber.Ctx) error {
	todoId, err := strconv.ParseInt(c.Params(params.TodoId), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusBadRequest),
			Message: message.InvalidId(domain.Model),
			Data: struct{}{},
		})
	}

	req := domain.TodoUpdateRequest{
		ID: todoId,
	}
	c.BodyParser(&req)

	if (
		req.Title == constant.EmptyString && 
		req.IsActive == nil &&
		req.Priority == constant.EmptyString) {

		return c.Status(http.StatusBadRequest).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusBadRequest), 
			Message: message.InvalidRequestBody,
			Data: struct{}{},
		})
	}

	if req.Priority != constant.EmptyString {
		if !slice.Includes(domain.PriorityAllList, req.Priority) {
			return c.Status(http.StatusBadRequest).JSON(web.BaseResponse{
				Status: http.StatusText(http.StatusBadRequest), 
				Message: message.ShoudMatchEnum(field.Priority, domain.PriorityAllList),
				Data: struct{}{},
			})
		}
	}

	res, err := h.Usecase.Update(c, req)
	if err != nil {
		return nil
	}

	return c.JSON(web.BaseResponse{
		Status: message.Success,
		Message: message.Success,
		Data: res,
	})
}

func (h *RESTHandler) Delete(c *fiber.Ctx) error {
	todoId, err := strconv.ParseInt(c.Params(params.TodoId), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusBadRequest),
			Message: message.InvalidId(domain.Model),
			Data: struct{}{},
		})
	}

	req := domain.TodoDeleteRequest{
		ID: todoId,
	}

	res, err := h.Usecase.Delete(c, req)
	if err != nil {
		return nil
	}

	return c.JSON(web.BaseResponse{
		Status: message.Success,
		Message: message.Success,
		Data: res,
	})
}

func (h *RESTHandler) GetAll(c *fiber.Ctx) error {
	req := domain.TodoGetAllRequest{		
		ActivityGroupID: int64(c.QueryInt(query.ActivityGroupID)),
	}

	res, err := h.Usecase.GetAll(c, req)
	if err != nil {
		return nil
	}

	return c.JSON(web.BaseResponse{
		Status: message.Success,
		Message: message.Success,
		Data: res,
	})
}

func (h *RESTHandler) GetOne(c *fiber.Ctx) error {
	todoId, err := strconv.ParseInt(c.Params(params.TodoId), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusBadRequest),
			Message: message.InvalidId(domain.Model),
			Data: struct{}{},
		})
	}

	req := domain.TodoGetOneRequest{
		ID: todoId,
	}

	res, err := h.Usecase.GetOne(c, req)
	if err != nil {
		return nil
	}

	return c.JSON(web.BaseResponse{
		Status: message.Success,
		Message: message.Success,
		Data: res,
	})
}