package repository

import (
	"github.com/p2-graded-challenge-2-jspheykel/internal/entity"
	"gorm.io/gorm"
)

type StudentRepo struct{ db *gorm.DB }

func NewStudentRepo(db *gorm.DB) *StudentRepo { return &StudentRepo{db: db} }

func (r *StudentRepo) Create(s *entity.Student) error { return r.db.Create(s).Error }

func (r *StudentRepo) FindByEmail(email string) (*entity.Student, error) {
	var s entity.Student
	if err := r.db.Where("email = ?", email).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *StudentRepo) GetByID(id int) (*entity.Student, error) {
	var s entity.Student
	if err := r.db.First(&s, "student_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *StudentRepo) GetMeWithEnrollments(id int) (*entity.Student, []entity.Enrollment, error) {
	var s entity.Student
	if err := r.db.First(&s, "student_id = ?", id).Error; err != nil {
		return nil, nil, err
	}
	var es []entity.Enrollment
	if err := r.db.
		Preload("Course").
		Where("student_id = ? AND deleted_at IS NULL", id).
		Find(&es).Error; err != nil {
		return &s, nil, err
	}
	return &s, es, nil
}
