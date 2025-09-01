package service

import (
	"errors"

	"github.com/p2-graded-challenge-2-jspheykel/internal/repository"
)

type EnrollmentService struct {
	courses     *repository.CourseRepo
	enrollments *repository.EnrollmentRepo
}

func NewEnrollmentService(c *repository.CourseRepo, e *repository.EnrollmentRepo) *EnrollmentService {
	return &EnrollmentService{courses: c, enrollments: e}
}

func (s *EnrollmentService) Enroll(studentID, courseID int) (any, error) {
	// ensure course exists
	if _, err := s.courses.Get(courseID); err != nil {
		return nil, errors.New("course not found")
	}
	return s.enrollments.Create(studentID, courseID)
}

func (s *EnrollmentService) Unenroll(studentID, enrollmentID int) (any, error) {
	return s.enrollments.SoftDelete(enrollmentID, studentID)
}
