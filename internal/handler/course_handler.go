package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/p2-graded-challenge-2-jspheykel/internal/service"
)

// @Summary List courses
// @Security Bearer
// @Tags Courses
// @Produce json
// @Success 200 {array} entity.Course
// @Router /courses [get]
type CourseHandler struct{ svc *service.CourseService }

func NewCourseHandler(s *service.CourseService) *CourseHandler { return &CourseHandler{svc: s} }

func (h *CourseHandler) List(c echo.Context) error {
	cs, err := h.svc.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to list courses"})
	}
	return c.JSON(http.StatusOK, cs)
}
