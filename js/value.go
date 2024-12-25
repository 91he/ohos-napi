package js

import napi "github.com/likuai2010/ohos-napi"

type Value struct {
	Env   Env
	Value napi.Value
}

func (v Value) GetEnv() Env {
	return v.Env
}
