package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gitub.com/umardev500/gopos/internal/app/contract"
	"gitub.com/umardev500/gopos/internal/app/models"
	"gitub.com/umardev500/gopos/pkg/auth"
	pkgUtil "gitub.com/umardev500/gopos/pkg/util"
	"gitub.com/umardev500/gopos/pkg/validator"
)

type authService struct {
	repo     contract.AuthRepository
	validate validator.Validator
}

func NewAuthService(repo contract.AuthRepository, v validator.Validator) contract.AuthService {
	return &authService{
		repo:     repo,
		validate: v,
	}
}

func (s *authService) Login(ctx context.Context, payload *models.LoginRequest) *pkgUtil.Response {
	// Validate input payload
	res := s.validate.Struct(payload).Response()
	if res != nil {
		return res
	}

	// Fetch user details from repository
	user, err := s.repo.Login(ctx, payload.Username)
	if err != nil {
		return &pkgUtil.Response{}
	}

	// Verify password
	ok := pkgUtil.CheckPasswordHash(payload.Password, user.PasswordHash)
	if !ok {
		return &pkgUtil.Response{
			Message: "Invalid username or password",
		}
	}

	// Generate JWT token
	token, err := s.generateJWTToken(user)
	if err != nil {
		return &pkgUtil.Response{}
	}

	return &pkgUtil.Response{
		Success: true,
		Message: "Login success",
		Data: map[string]string{
			"token": token,
		},
	}
}

// generateJWTToken creates a JWT token for the authenticated user.
func (s *authService) generateJWTToken(user *models.AuthUser) (string, error) {
	claims := jwt.MapClaims{
		"uid": user.ID,
		"tid": user.TenantID,
		"bid": user.BranchID,
		"exp": time.Now().UTC().Add(time.Hour * 1).Unix(),
		"iat": time.Now().UTC().Unix(),
	}
	return auth.GenerateJWT(claims)
}
