package napi

/*
#cgo CFLAGS: -DDEBUG
#cgo CFLAGS: -D_DEBUG
#cgo CFLAGS: -DV8_ENABLE_CHECKS
#cgo CFLAGS: -DNAPI_EXPERIMENTAL
#cgo CFLAGS: -I/usr/include/napi
#cgo CXXFLAGS: -std=c++11

#cgo linux LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo LDFLAGS: -lace_napi.z
#cgo LDFLAGS: -stdlib=libc++
*/
import "C"
