package main

import (
	napi "github.com/91he/ohos-napi"
	"github.com/91he/ohos-napi/entry"
)

func init() {
	entry.Export("getCallback", GetCallbackHandler)
}

func GetCallbackHandler(env napi.Env, info napi.CallbackInfo) napi.Value {
	result, _ := napi.CreateFunction(
		env,
		"callback",
		func(env napi.Env, info napi.CallbackInfo) napi.Value {
			result, _ := napi.CreateStringUtf8(env, "hello world")
			return result
		},
	)

	return result
}

func main() {}
