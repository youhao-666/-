package handler

import (
	"github.com/gin-gonic/gin"
	"/internal/service"
)

type StudentHandler struct {
	*Handler
	studentService service.StudentService
}

func NewStudentHandler(
    handler *Handler,
    studentService service.StudentService,
) *StudentHandler {
	return &StudentHandler{
		Handler:      handler,
		studentService: studentService,
	}
}

func (h *StudentHandler) GetStudent(ctx *gin.Context) {

}
