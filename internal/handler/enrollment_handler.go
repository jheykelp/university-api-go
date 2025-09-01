package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/p2-graded-challenge-2-jspheykel/internal/service"
)

// @Summary Enroll into a course
// @Security Bearer
// @Tags Enrollments
// @Accept json
// @Produce json
// @Param body body enrollRequest true "Enroll request"
// @Success 201 {object} enrollmentResponse
// @Failure 400 {object} errorResponse
// @Router /enrollments [post]
type EnrollmentHandler struct{ svc *service.EnrollmentService }

func NewEnrollmentHandler(s *service.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{svc: s}
}

func (h *EnrollmentHandler) Enroll(c echo.Context) error {
	var p struct {
		CourseID int `json:"course_id"`
	}
	if err := c.Bind(&p); err != nil || p.CourseID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "course_id required"})
	}
	userID := c.Get("user_id").(int)
	en, err := h.svc.Enroll(userID, p.CourseID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, en)
}

// @Summary Delete (soft) an enrollment
// @Security Bearer
// @Tags Enrollments
// @Param id path int true "Enrollment ID"
// @Success 200 {object} enrollmentResponse
// @Router /enrollments/{id} [delete]
func (h *EnrollmentHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid id"})
	}
	userID := c.Get("user_id").(int)
	en, err := h.svc.Unenroll(userID, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, en)
}
