package proto

import "context"

type IService interface {
	GenerateToken(context.Context, *GenerateTokenReqeust) (*GenerateTokenResponse, error)
}
