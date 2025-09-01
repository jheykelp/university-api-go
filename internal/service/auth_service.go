package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/p2-graded-challenge-2-jspheykel/internal/entity"
	"github.com/p2-graded-challenge-2-jspheykel/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	students  *repository.StudentRepo
	jwtSecret []byte
}

func NewAuthService(st *repository.StudentRepo, secret string) *AuthService {
	return &AuthService{students: st, jwtSecret: []byte(secret)}
}

type RegisterInput struct {
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	Address     string    `json:"address" validate:"required"`
	Password    string    `json:"password" validate:"required,min=6"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
}

func (s *AuthService) Register(in RegisterInput) (*entity.Student, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	st := &entity.Student{
		FirstName:    in.FirstName,
		LastName:     in.LastName,
		Email:        in.Email,
		Address:      in.Address,
		DateOfBirth:  in.DateOfBirth,
		PasswordHash: string(hash),
	}
	if err := s.students.Create(st); err != nil {
		return nil, err
	}
	// do NOT expose password hash in response (spec requires that)  [oai_citation:6‡README.md](file-service://file-TNPfMDEXjpad9iQmM9NYXc)
	st.PasswordHash = ""
	return st, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	st, err := s.students.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}
	if bcrypt.CompareHashAndPassword([]byte(st.PasswordHash), []byte(password)) != nil {
		return "", errors.New("invalid email or password")
	}
	claims := jwt.MapClaims{
		"sub": st.StudentID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tok.SignedString(s.jwtSecret)
}
