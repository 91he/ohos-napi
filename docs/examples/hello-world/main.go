package main

import (
	"fmt"

	napi "github.com/91he/ohos-napi"
	"github.com/91he/ohos-napi/entry"
)

func init() {
	entry.Export("hello", HelloHandler)
}

func HelloHandler(env napi.Env, info napi.CallbackInfo) napi.Value {
	fmt.Println("hello world!")
	return nil
}

func main() {}
