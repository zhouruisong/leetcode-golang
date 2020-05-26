package main

import (
	"fmt"
	"unsafe"
)

/*
type StringHeader struct {
	Data uintptr
	Len  int
}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

在 golang 中，遇到不需要解析的 json 数据，可以将其类型声明为json.RawMessage
尽量避免传入参数为interface类型,无法在编译时确定具体类型的

func readJsonReq(data []byte, req *model.GatewayReqBody) error {
err := jsoniter.Unmarshal(data, req)
...
}

优化为:
func readJsonReq(data []byte, req *model.GatewayReqBody) error {
	var tmp model.GatewayReqBody
	err := jsoniter.Unmarshal(data, &tmp)
	...
}

type GatewayReqBody struct {
 Header  GatewayReqBodyHeader `json:"header"`
 Payload json.RawMessage      `json:"payload"`
}
*/

//压测优化string 和 byte的转换拷贝
func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Str2Bytes(s string) []byte {
	p := &s
	fmt.Printf("%+v\n", p)

	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func main() {
	x := "0123"
	fmt.Println(Str2Bytes(x))
}
