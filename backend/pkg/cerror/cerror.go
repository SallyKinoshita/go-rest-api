package cerror

import (
	"bytes"
	"errors"
	"io"
	"runtime"
	"slices"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CError struct {
	err     error
	details []Detail
	frames  *runtime.Frames
}

func New(developerMessage string, detail Detail) *CError {
	detail.validateAndSetDetailInternal()

	return newCError(errors.New(developerMessage), detail, newFrames())
}

func Wrap(err error, detail Detail) *CError {
	if cerr := ConvertCError(err); cerr != nil {
		return newCError(err, detail, cerr.frames)
	}

	return newCError(err, detail, newFrames())
}

func newCError(err error, detail Detail, frames *runtime.Frames) *CError {
	return &CError{
		err:     err,
		details: []Detail{detail},
		frames:  frames,
	}
}

func newFrames() *runtime.Frames {
	pc := make([]uintptr, 64)
	n := runtime.Callers(3, pc)
	pc = pc[:n]
	return runtime.CallersFrames(pc)
}

func (c *CError) Error() string {
	if c == nil || c.err == nil {
		return "unknown error"
	}
	return c.err.Error()
}

func (c *CError) StatusCode() int {
	if c == nil || len(c.details) == 0 {
		return StatusInternalServerError.code()
	}
	return c.details[0].status.code()
}

func (c *CError) StatusText() string {
	if c == nil || len(c.details) == 0 {
		return StatusInternalServerError.text()
	}
	return c.details[0].status.text()
}

func (c *CError) ContainsStatus(statuses ...status) bool {
	if c == nil {
		return false
	}

	return slices.Contains(statuses, c.details[0].status)
}

func (c *CError) StackTrace() string {
	if c == nil {
		return ""
	}
	if c.frames == nil {
		return ""
	}

	var buf bytes.Buffer

	for {
		frame, more := c.frames.Next()

		io.WriteString(&buf, frame.Function)
		io.WriteString(&buf, "\n\t")
		io.WriteString(&buf, frame.File)
		io.WriteString(&buf, ":")
		io.WriteString(&buf, strconv.Itoa(frame.Line))

		if !more {
			break
		}
	}

	return buf.String()
}

func (c *CError) ConvertEchoError() echo.Map {
	if c == nil {
		return echo.Map{
			"status": StatusInternalServerError.text(),
			"detail": messageInternal,
		}
	}

	return echo.Map{
		"status": c.StatusText(),
		"detail": c.details[0].clientMessage,
	}
}

func ConvertCError(err error) *CError {
	var cerr *CError
	if yes := errors.As(err, &cerr); !yes {
		return nil
	}

	return cerr
}
