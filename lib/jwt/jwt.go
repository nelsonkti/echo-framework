package jwt

import (
    "errors"
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "echo-framework/config"
)

type TokenData struct {
    StaffId  uint32
    ExpireAt int64
}

func ParseToken(tokenString string) (*TokenData, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Don't forget to validate the alg is what you expect:
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        // hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
        return []byte(config.AppSign), nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if ok == false || token.Valid == false {
        return nil, errors.New("validation failed")
    }

    return &TokenData{StaffId: uint32(claims["staff_id"].(float64)), ExpireAt: int64(claims["expire_at"].(float64))}, nil
}

func MakeToken(data TokenData) string {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "staff_id":  data.StaffId,
        "expire_at": data.ExpireAt,
    })

    tokenString, _ := token.SignedString([]byte(config.AppSign))
    return tokenString
}
