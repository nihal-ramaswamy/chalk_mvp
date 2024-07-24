package constants

import "time"

const (
	WRITE_TIMEOUT    = 10 * time.Second
	PING_TIMEOUT     = time.Minute
	PONG_TIMEOUT     = time.Minute * 9 / 10
	MAX_MESSAGE_SIZE = 512

	READ_BUFFER_SIZE  = 1024
	WRITE_BUFFER_SIZE = 1024
)
