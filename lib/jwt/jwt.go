package jwt

import (
    "echo-framework/config"
    "echo-framework/lib/helper"
    "errors"
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "reflect"
)

type TokenData struct {
    StaffId  uint64
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

    var staffId uint64
    if reflect.TypeOf(claims["user_id"]).String() == "string" {
        staffId = helper.StrToUint64(fmt.Sprintf("%s", claims["user_id"]))
    } else {
        staffId = uint64(claims["user_id"].(float64))
    }

    return &TokenData{StaffId: staffId, ExpireAt: int64(claims["exp"].(float64))}, nil

}

func MakeToken(data TokenData) string {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "staff_id":  data.StaffId,
        "expire_at": data.ExpireAt,
    })

    tokenString, _ := token.SignedString([]byte(config.AppSign))
    return tokenString
}

