package handler

import (
    "time"

    "github.com/p2-graded-challenge-2-jspheykel/internal/entity"
)

// Request/response DTOs for Swagger documentation.

type registerRequest struct {
    FirstName   string `json:"first_name"`
    LastName    string `json:"last_name"`
    Email       string `json:"email"`
    Address     string `json:"address"`
    Password    string `json:"password"`
    DateOfBirth string `json:"date_of_birth"` // YYYY-MM-DD
}

type registerResponse struct {
    Message string        `json:"message"`
    Data    entity.Student `json:"data"`
}

type errorResponse struct {
    Message string `json:"message"`
}

type loginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type tokenResponse struct {
    Token string `json:"token"`
}

type meEnrollment struct {
    ID             int       `json:"id"`
    CourseID       int       `json:"course_id"`
    CourseName     string    `json:"course_name"`
    EnrollmentDate time.Time `json:"enrollment_date"`
}

type meResponse struct {
    ID           int            `json:"id"`
    FirstName    string         `json:"first_name"`
    LastName     string         `json:"last_name"`
    Email        string         `json:"email"`
    Address      string         `json:"address"`
    DateOfBirth  string         `json:"date_of_birth"`
    Enrollments  []meEnrollment `json:"enrollments"`
}

type enrollRequest struct {
    CourseID int `json:"course_id"`
}

type enrollmentResponse struct {
    ID             int         `json:"id"`
    StudentID      int         `json:"student_id"`
    CourseID       int         `json:"course_id"`
    EnrollmentDate time.Time   `json:"enrollment_date"`
    DeletedAt      *time.Time  `json:"deleted_at,omitempty"`
    Course         entity.Course `json:"course"`
}

