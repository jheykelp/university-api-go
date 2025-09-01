package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/p2-graded-challenge-2-jspheykel/internal/repository"
	"github.com/p2-graded-challenge-2-jspheykel/internal/service"
)

type StudentHandler struct {
	auth *service.AuthService
	st   *repository.StudentRepo
}

// @Summary Register new student
// @Tags Students
// @Accept json
// @Produce json
// @Param body body registerRequest true "Student payload"
// @Success 201 {object} registerResponse
// @Failure 400 {object} errorResponse
// @Router /students/register [post]
func (h *StudentHandler) Register(c echo.Context) error {
	type payload struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		Address     string `json:"address"`
		Password    string `json:"password"`
		DateOfBirth string `json:"date_of_birth"` // "2003-05-20"
	}
	var p payload
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid payload"})
	}
	dob, err := time.Parse("2006-01-02", p.DateOfBirth)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid date_of_birth (YYYY-MM-DD)"})
	}
	st, err := h.auth.Register(service.RegisterInput{
		FirstName: p.FirstName, LastName: p.LastName, Email: p.Email,
		Address: p.Address, Password: p.Password, DateOfBirth: dob,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "student registered",
		"data":    st,
	})
}

// @Summary Login and get JWT
// @Tags Students
// @Accept json
// @Produce json
// @Param body body loginRequest true "Login payload"
// @Success 200 {object} tokenResponse
// @Failure 401 {object} errorResponse
// @Router /students/login [post]
func (h *StudentHandler) Login(c echo.Context) error {
	var p struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid payload"})
	}
	tok, err := h.auth.Login(p.Email, p.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"token": tok})
}

// @Summary Current student profile
// @Security Bearer
// @Tags Students
// @Produce json
// @Success 200 {object} meResponse
// @Router /students/me [get]
func (h *StudentHandler) Me(c echo.Context) error {
	userID := c.Get("user_id").(int)
	st, ens, err := h.st.GetMeWithEnrollments(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to load profile"})
	}
	type enrollOut struct {
		ID             int       `json:"id"`
		CourseID       int       `json:"course_id"`
		CourseName     string    `json:"course_name"`
		EnrollmentDate time.Time `json:"enrollment_date"`
	}
	out := make([]enrollOut, 0, len(ens))
	for _, e := range ens {
		out = append(out, enrollOut{ID: e.EnrollmentID, CourseID: e.CourseID, CourseName: e.Course.Name, EnrollmentDate: e.EnrollmentDate})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"id":            st.StudentID,
		"first_name":    st.FirstName,
		"last_name":     st.LastName,
		"email":         st.Email,
		"address":       st.Address,
		"date_of_birth": st.DateOfBirth.Format("2006-01-02"),
		"enrollments":   out,
	})
}

func NewStudentHandler(auth *service.AuthService, st *repository.StudentRepo) *StudentHandler {
	return &StudentHandler{auth: auth, st: st}
}
