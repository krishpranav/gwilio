/**
 * @file sanity.go
 * @author Krisna Pranav
 * @version 1.0
 * @date 2025-05-09
 *
 * @copyright Copyright (c) 2025 Krisna Pranav
 *
 */

package helper

import "regexp"

func IsValidRecordingStatusCallbackEvent(event string) bool {
	switch event {
		case 
			"in-progress",
			"completed",
			"absent":
			return true
	}

	return false
}

func IsValidTrim(trim string) bool {
	switch trim {
		case 
			"trim-silence":
			"do-not-trim":
			return true
	}

	return false
}

func IsValidCallReason(reason string) bool {
	return len(reason) <= 50
} 

func DtmfSanity(number string) string {
	reg, err := regexp.Compile("")

	if err != nil {
		return number
	}

	return reg.ReplaceAllString(number, "")
}