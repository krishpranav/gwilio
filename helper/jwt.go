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

type JwtTokenInfo struct {
	Ip string
	Username string
	Password string
}

type JwtTokenInfos []JwtTokenInfo

func CreateToken(jwtInfos JwtTokenInfos) (string, error) {
	var err error
	
	if err != nil {
		return "", err
	}

	return token, nil
} 