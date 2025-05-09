/**
 * @file error.go
 * @author Krisna Pranav
 * @version 1.0
 * @date 2025-05-09
 *
 * @copyright Copyright (c) 2025 Krisna Pranav
 *
 */

package constant

import "errors"

var (
	ErrEmptyXML = errors.New("XML Repsonse is empty")
	ErrInvalidXML = errors.New("XML Response is invalid")
	ErrEmptyRedirect = errors.New("XML Redirect is empty")
	ErrGatherTimeout = errors.New("gather Timeout")
)