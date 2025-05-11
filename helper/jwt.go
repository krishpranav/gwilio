/**
 * @file jwt.go
 * @author Krisna Pranav
 * @version 1.0
 * @date 2025-05-09
 *
 * @copyright Copyright (c) 2025 Krisna Pranav
 *
 */

package helper

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtTokenInfo struct {
	Ip string
	Username string
	Password string
}

type JwtTokenInfos []JwtTokenInfo

func CreateToken(jwtInfos JwtTokenInfos) (string, error) {
	var err error
	
	os.Setenv("ACCESS_SECRET", "RZ+w1Vr/dk4nZHvd/B7av/pOGiNzYlPZ") 
	atClaims := jwt.MapClaims{}

	for _, jwtInfo := range jwtInfos {
		atClaims[jwtInfo.Ip] = jwtInfo.Username
		atClaims[jwtInfo.Username] = jwtInfo.Password
	}

	atClaims["authorized"] = true
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}