package repository

import (
	"github.com/p2-graded-challenge-2-jspheykel/internal/entity"
	"gorm.io/gorm"
)

type CourseRepo struct{ db *gorm.DB }

func NewCourseRepo(db *gorm.DB) *CourseRepo { return &CourseRepo{db: db} }

func (r *CourseRepo) List() ([]entity.Course, error) {
	var cs []entity.Course
	return cs, r.db.Find(&cs).Error
}

func (r *CourseRepo) Get(id int) (*entity.Course, error) {
	var c entity.Course
	if err := r.db.First(&c, "course_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}
