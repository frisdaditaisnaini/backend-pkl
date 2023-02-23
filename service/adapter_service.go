package service

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/adapter"
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/config"

	"github.com/golang-jwt/jwt"
)

type svc struct {
	c    config.Config
	repo adapter.AdapterRepository
}

func (s *svc) ClaimToken(bearer *jwt.Token) string {
	claim := bearer.Claims.(jwt.MapClaims)
	username := fmt.Sprintf("%v", claim["username"])

	return username
}

func NewService(repo adapter.AdapterRepository, c config.Config) adapter.AdapterService {
	return &svc{
		repo: repo,
		c:    c,
	}
}
