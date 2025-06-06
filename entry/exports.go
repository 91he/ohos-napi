package entry

import (
	napi "github.com/91he/ohos-napi"
)

type napiGoExport struct {
	Name     string
	Callback napi.Callback
}

var napiGoGlobalExports []napiGoExport

func Export(name string, callback napi.Callback) {
	napiGoGlobalExports = append(napiGoGlobalExports, napiGoExport{
		Name:     name,
		Callback: callback,
	})
}
