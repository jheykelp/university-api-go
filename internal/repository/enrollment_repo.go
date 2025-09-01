package repository

import (
	"time"

	"github.com/p2-graded-challenge-2-jspheykel/internal/entity"
	"gorm.io/gorm"
)

type EnrollmentRepo struct{ db *gorm.DB }

func NewEnrollmentRepo(db *gorm.DB) *EnrollmentRepo { return &EnrollmentRepo{db: db} }

func (r *EnrollmentRepo) Create(studentID, courseID int) (*entity.Enrollment, error) {
	en := &entity.Enrollment{
		StudentID:      studentID,
		CourseID:       courseID,
		EnrollmentDate: time.Now().UTC(),
	}
	// Unique index on (student_id, course_id) WHERE deleted_at IS NULL will prevent duplicates.
	if err := r.db.Create(en).Error; err != nil {
		return nil, err
	}
	_ = r.db.Preload("Course").First(en, "enrollment_id = ?", en.EnrollmentID)
	return en, nil
}

func (r *EnrollmentRepo) SoftDelete(id, studentID int) (*entity.Enrollment, error) {
	var en entity.Enrollment
	if err := r.db.Preload("Course").
		First(&en, "enrollment_id = ? AND student_id = ? AND deleted_at IS NULL", id, studentID).Error; err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	if err := r.db.Model(&en).Update("deleted_at", &now).Error; err != nil {
		return nil, err
	}
	return &en, nil
}
