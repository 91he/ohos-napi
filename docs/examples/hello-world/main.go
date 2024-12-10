package main

import (
	"fmt"

	"github.com/akshayganeshen/napi-go/entry"
	napi "github.com/likuai2010/ohos-napi"
)

func init() {
	entry.Export("hello", HelloHandler)
}

func HelloHandler(env napi.Env, info napi.CallbackInfo) napi.Value {
	fmt.Println("hello world!")
	return nil
}

func main() {}
