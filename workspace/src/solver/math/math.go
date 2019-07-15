package math

import "C"
import (
	"log"
	"syscall"
)

const path = "G:/test/test/math_x64.dll" // Define the path to the file math_x64.dll or math_x86.dll

func Div(arg1 int, arg2 int, id int, ch chan int, chi chan int) {
	h, err := syscall.LoadLibrary(path)
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.FreeLibrary(h)
	proc, err := syscall.GetProcAddress(h, "Div")
	if err != nil {
		log.Fatal(err)
	}
	n, _, err := syscall.Syscall(uintptr(proc), 0, uintptr(arg1), uintptr(arg2), 0)
	ch <- int(n)
	chi <- id
}