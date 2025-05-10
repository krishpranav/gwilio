/**
 * @file heartbeat.go
 * @author Krisna Pranav
 * @version 1.0
 * @date 2025-05-09
 *
 * @copyright Copyright (c) 2025 Krisna Pranav
 *
 */

package models

type HeartBeatPrivRequest struct {
	AccountID string  `json:"accountid"`
	CallID    string  `json:"callid"`
	Rate      float64 `json:"rate"`
	Pulse     int64   `json:"pulse"`
	Duration  int64   `json:"duration"`
}