package service

import (
	"github.com/p2-graded-challenge-2-jspheykel/internal/entity"
	"github.com/p2-graded-challenge-2-jspheykel/internal/repository"
)

type CourseService struct{ repo *repository.CourseRepo }

func NewCourseService(r *repository.CourseRepo) *CourseService { return &CourseService{repo: r} }

func (s *CourseService) List() ([]entity.Course, error) { return s.repo.List() }
