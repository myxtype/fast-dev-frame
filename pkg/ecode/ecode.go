package ecode

import (
	"fmt"
	"strings"
)

var (
	codes = map[int]struct{}{}
)

// New Error
func New(code int, msg string) Error {
	if code < 1000 {
		panic("error code must be greater than 1000")
	}
	return add(code, msg)
}

// add only inner error
func add(code int, msg string) Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", code))
	}
	codes[code] = struct{}{}
	return Error{
		code: code, message: msg,
	}
}

type Errors interface {
	// Error return Code in string form
	Error() string
	// Code get error code.
	Code() int
	// Message get code message.
	Message() string
	// Reload Message
	Reload(string) Error
}

type Error struct {
	code    int
	message string
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Code() int {
	return e.code
}

func (e Error) Message() string {
	return e.message
}

func (e Error) Reload(message string) Error {
	e.message = message
	return e
}

func String(msg string) Error {
	if msg == "" || strings.ToLower(msg) == "ok" {
		return Ok
	}
	return Error{
		code:    errServerCode,
		message: msg,
	}
}

func Cause(err error) Errors {
	if err == nil {
		return Ok
	}
	cause, ok := err.(Errors)
	if ok {
		return cause
	}
	return String(err.Error())
}
