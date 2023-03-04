package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fahmiaz411/devcode/helper/constant"
	"github.com/fahmiaz411/devcode/helper/field"
	"github.com/fahmiaz411/devcode/helper/message"
	"github.com/fahmiaz411/devcode/helper/params"
	"github.com/fahmiaz411/devcode/helper/web"
	"github.com/fahmiaz411/devcode/modules/activity/domain"
	"github.com/fahmiaz411/devcode/modules/activity/interfaces"

	"github.com/gofiber/fiber/v2"
)

type RESTHandler struct {
	Usecase interfaces.ActivityUsecase
}

func NewRESTHandler(f fiber.Router, usecase interfaces.ActivityUsecase) {
	handler := &RESTHandler{
		Usecase: usecase,
	}

	f.Post("/activity-groups", handler.Create)

	f.Patch(fmt.Sprintf("/activity-groups/:%s", params.ActivityId), handler.Update)

	f.Delete(fmt.Sprintf("/activity-groups/:%s", params.ActivityId), handler.Delete)

	f.Get("/activity-groups", handler.GetAll)

	f.Get(fmt.Sprintf("/activity-groups/:%s", params.ActivityId), handler.GetOne)
}

func (h *RESTHandler) Create(c *fiber.Ctx) error {
	req := domain.ActivityCreateRequest{}
	c.BodyParser(&req)

	if req.Title == constant.EmptyString {
		return c.Status(http.StatusBadRequest).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusBadRequest),
			Message: message.CanotNull(field.Title),
			Data: struct{}{},
		})
	}

	res, err := h.Usecase.Create(c, req)
	if err != nil {
		return nil
	}

	return c.Status(http.StatusCreated).JSON(web.BaseResponse{
		Status: message.Success,
		Message: message.Success,
		Data: res,
	})
}

func (h *RESTHandler) Update(c *fiber.Ctx) error {
	activityId, err := strconv.ParseInt(c.Params(params.ActivityId), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusBadRequest),
			Message: message.InvalidId(domain.Model),
			Data: struct{}{},
		})
	}

	req := domain.ActivityUpdateRequest{
		ID: activityId,
	}
	c.BodyParser(&req)

	if req.Title == constant.EmptyString {
		return c.Status(http.StatusBadRequest).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusBadRequest),
			Message: message.CanotNull(field.Title),
			Data: struct{}{},
		})
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
	activityId, err := strconv.ParseInt(c.Params(params.ActivityId), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusBadRequest),
			Message: message.InvalidId(domain.Model),
			Data: struct{}{},
		})
	}

	req := domain.ActivityDeleteRequest{
		ID: activityId,
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
	req := domain.ActivityGetAllRequest{		
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
	activityId, err := strconv.ParseInt(c.Params(params.ActivityId), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.BaseResponse{
			Status: http.StatusText(http.StatusBadRequest),
			Message: message.InvalidId(domain.Model),
			Data: struct{}{},
		})
	}

	req := domain.ActivityGetOneRequest{
		ID: activityId,
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