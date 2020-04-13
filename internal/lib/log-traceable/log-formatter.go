package log_traceable

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"runtime"
)

func GetMessage(ctx context.Context, message ...interface{}) (rString string) {
	// file depth 1 read the first method locations where the external function
	// call the the location.
	fileDepth := 1
	_, file, line, _ := runtime.Caller(fileDepth)

	// message array from the message append to the one string
	msgStr := ""
	if len(message) > 0 {
		for _, msg := range message {
			msgStr = fmt.Sprintf("%v %v", msgStr, msg)
		}
	}
	// get the uuid from the context and format the string according to the application
	// and format the log message
	// format : <uuid><tab><file>:<line number><space><message 1>....<space><message n>
	uuidInf := ctx.Value("uuid_str")
	rString = fmt.Sprintf("%v\t%v:%v %v", uuid.New().String(), file, line, msgStr)
	if uuidInf != nil {
		rString = fmt.Sprintf("%v\t%v:%v %v", uuidInf, file, line, msgStr)
	}
	return
}
