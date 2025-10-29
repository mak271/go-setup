package jwt

import "github.com/golang-jwt/jwt/v5"

type JWTData struct {
	Email string
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{Secret: secret}
}

func (j *JWT) Create(data JWTData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})
	signedStr, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return signedStr, nil
}

func (j *JWT) Parse(token string) (bool, *JWTData) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	email := parsedToken.Claims.(jwt.MapClaims)["email"]
	return parsedToken.Valid, &JWTData{
		Email: email.(string),
	}
}
