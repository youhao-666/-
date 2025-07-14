package handler

import (
	"github.com/gin-gonic/gin"
	"server_go/internal/service"
)

type PeopleHandler struct {
	*Handler
	peopleService service.PeopleService
}

func NewPeopleHandler(
    handler *Handler,
    peopleService service.PeopleService,
) *PeopleHandler {
	return &PeopleHandler{
		Handler:      handler,
		peopleService: peopleService,
	}
}

func (h *PeopleHandler) GetPeople(ctx *gin.Context) {

}
