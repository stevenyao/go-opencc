package opencc

import "unsafe"

// #cgo LDFLAGS: -lopencc
// #include "opencc.h"
import "C"

type Converter struct {
	id unsafe.Pointer
}

func NewConverter(config string) *Converter {
	c := Converter{}
	c.id = C.Opencc_New(C.CString(config));
	return &c
}

func (c *Converter)Convert(input string) string {
	output := C.Opencc_Convert(c.id, C.CString(input))
	defer C.Opencc_Free_String(output)
	return C.GoString(output)
}

func (c *Converter)Close(){
	C.Opencc_Delete(c.id)
}

func Convert(input string, config string) string {
	c := NewConverter(config)
	defer c.Close()
	return c.Convert(input)
}

func ConvertAsync(input string, config string, callback func(output string)) {
	go func() {
		callback(Convert(input, config))
	}()
}