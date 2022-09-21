package token

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTTokenGen struct {
	issuer     string
	nowFunc    func() time.Time
	privateKey *rsa.PrivateKey
}

func NewJWTTokenGen(issuer string, privateKey *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{
		issuer:     issuer,
		privateKey: privateKey,
		nowFunc:    time.Now,
	}
}

// GeneRateToken generates a token.
func (t *JWTTokenGen) GenerateToken(accountId string, exp time.Duration) (string, error) {
	nowTime := t.nowFunc().Unix()
	tnk := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Issuer:    t.issuer,
		IssuedAt:  nowTime,
		ExpiresAt: nowTime + int64(exp.Seconds()),
		Subject:   accountId,
	})
	return tnk.SignedString(t.privateKey)
}
