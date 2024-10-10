package telefleet

import (
	"reflect"
	"runtime"
	"strings"

	"gopkg.in/telebot.v3"
)

// getFuncName is a helper function that returns the name of a function.
// It takes a middleware function and an interface{} as arguments.
// If the interface{} is not nil, it uses the interface{} to get the function name.
// Otherwise, it uses the middleware function to get the function name.
func getFuncName(middlewareFunct telebot.MiddlewareFunc, function interface{}) string {
	var funcName string
	if function != nil {
		funcName = runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	} else {
		funcName = runtime.FuncForPC(reflect.ValueOf(middlewareFunct).Pointer()).Name()
	}
	// The function name returned by runtime.FuncForPC() is in the format of "package.function".
	// We only want the function name, so we extract it by finding the last occurrence of "." or "/" and taking the substring after it.
	lastSlash := strings.LastIndex(funcName, "/")
	lastDot := strings.LastIndex(funcName, ".")
	if lastSlash >= 0 && lastDot >= 0 && lastSlash < lastDot {
		funcName = funcName[lastSlash+1 : lastDot]
	}
	lastDot = strings.LastIndex(funcName, ".")
	if lastDot >= 0 {
		funcName = funcName[lastDot+1:]
	}
	return funcName
}
