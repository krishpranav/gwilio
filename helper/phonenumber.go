/**
 * @file phonenumber.go
 * @author Krisna Pranav
 * @version 1.0
 * @date 2025-05-09
 *
 * @copyright Copyright (c) 2025 Krisna Pranav
 *
 */

package helper

import "regexp"

func NumberSanity(number string) string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return number
	}
	return reg.ReplaceAllString(number, "")
}

func RemovePlus(number string) string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return number
	}
	return reg.ReplaceAllString(number, "")
}