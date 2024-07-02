package errs

import (
	"fmt"
	"runtime"
)

type error struct {
	info string
}

func new_error(info any) *error {
	callstack := make([]byte, 512)
	stackLen := runtime.Stack(callstack, false)
	callstack = callstack[2:stackLen]
	return &error{info: fmt.Sprintf("%v %v", info, callstack)}
}

func (e *error) Error() string {
	return e.info
}

type GinAppErr struct {
	*error
}

func NewGinAppErr(info any) *GinAppErr {
	return &GinAppErr{
		error: new_error(info),
	}
}

type FiberAppErr struct {
	*error
}

func NewFiberAppErr(info any) *FiberAppErr {
	return &FiberAppErr{
		error: new_error(info),
	}
}

type ChiAppErr struct {
	*error
}

func NewChiAppErr(info any) *ChiAppErr {
	return &ChiAppErr{
		error: new_error(info),
	}
}
