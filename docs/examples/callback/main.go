package main

import (
	"github.com/akshayganeshen/napi-go/entry"
	napi "github.com/likuai2010/ohos-napi"
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
