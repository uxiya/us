package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	RunCmd("clear")
	/*ch := make(chan string)
	go func() {
		for {
			select {
			case v := <-ch:
				arr := strings.Split(v, " ")
				if len(arr) > 0 {
					RunCmd(arr[0], arr[1:]...)
				}
			}
		}
	}()*/
	reader := bufio.NewReader(os.Stdin)
	for {
		bts, x, err := reader.ReadLine()
		if err == nil {
			arr := strings.Split(string(bts), " ")
			if len(arr) > 0 {
				RunCmd(arr[0], arr[1:]...)
			}
			reader.Discard(reader.Buffered())
		} else {
			fmt.Println(x, err)
		}

	}
}
func RunCmd(e string, args ...string) {
	cmd := exec.Command(e, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v ", "[zjt@u35s.com]")
	}
}
