package entity

import "time"

type Student struct {
	StudentID    int       `gorm:"primaryKey;column:student_id" json:"id"`
	FirstName    string    `gorm:"column:first_name" json:"first_name"`
	LastName     string    `gorm:"column:last_name" json:"last_name"`
	Email        string    `gorm:"uniqueIndex;column:email" json:"email"`
	Address      string    `gorm:"column:address" json:"address"`
	DateOfBirth  time.Time `gorm:"column:date_of_birth" json:"date_of_birth"`
	PasswordHash string    `gorm:"column:password_hash" json:"-"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Student) TableName() string { return "students" }

type Department struct {
	DepartmentID int    `gorm:"primaryKey;column:department_id" json:"department_id"`
	Name         string `gorm:"column:name" json:"name"`
	Description  string `gorm:"column:description" json:"description"`
}

func (Department) TableName() string { return "departments" }

type Course struct {
	CourseID     int    `gorm:"primaryKey;column:course_id" json:"id"`
	Name         string `gorm:"column:name" json:"name"`
	Description  string `gorm:"column:description" json:"description"`
	DepartmentID int    `gorm:"column:department_id" json:"department_id"`
	Credits      int16  `gorm:"column:credits" json:"credits"`
}

func (Course) TableName() string { return "courses" }

type Enrollment struct {
	EnrollmentID   int        `gorm:"primaryKey;column:enrollment_id" json:"id"`
	StudentID      int        `gorm:"column:student_id" json:"student_id"`
	CourseID       int        `gorm:"column:course_id" json:"course_id"`
	EnrollmentDate time.Time  `gorm:"column:enrollment_date" json:"enrollment_date"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Course         Course     `gorm:"foreignKey:CourseID" json:"course"`
}

func (Enrollment) TableName() string { return "enrollments" }

type Professor struct {
	ProfessorID int    `gorm:"primaryKey;column:professor_id" json:"professor_id"`
	FirstName   string `gorm:"column:first_name" json:"first_name"`
	LastName    string `gorm:"column:last_name" json:"last_name"`
	Email       string `gorm:"uniqueIndex;column:email" json:"email"`
	Address     string `gorm:"column:address" json:"address"`
}

func (Professor) TableName() string { return "professors" }
