package clog

type logType string

const (
	LogTypeApp   logType = "app"
	LogTypeBatch logType = "batch"
	LogTypeDB    logType = "db"
)
