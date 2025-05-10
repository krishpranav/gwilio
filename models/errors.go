/**
 * @file errors.go
 * @author Krisna Pranav
 * @version 1.0
 * @date 2025-05-09
 *
 * @copyright Copyright (c) 2025 Krisna Pranav
 *
 */

package models

import "net/http"

type RequestError struct {
	StatusCode int
	Err error
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func (r *RequestError) RatingRoutingMissing() bool {
	return r.StatusCode == http.StatusServiceUnavailable
}

func (r *RequestError) NestedDialElement() bool {
	return r.StatusCode == http.StatusMethodNotAllowed
}