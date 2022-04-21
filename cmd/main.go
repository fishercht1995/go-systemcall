package main

import (
	"syscall"
	"fmt"
	"os"
)

func main(){
	err := syscall.Exec("/bin/ls", []string{"-al"}, os.Environ())
	fmt.Println(err)
}
