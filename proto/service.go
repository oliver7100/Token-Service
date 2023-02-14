package proto

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
)

type service struct {
	UnimplementedAuthServiceServer
}

func (s *service) GenerateToken(ctx context.Context, req *GenerateTokenReqeust) (*GenerateTokenResponse, error) {
	claims := jwt.MapClaims{
		"username": req.Username,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	st, err := t.SignedString([]byte("secret"))

	if err != nil {
		return nil, err
	}

	return &GenerateTokenResponse{
		Token: st,
	}, nil
}

func NewService() *service {
	return &service{}
}
