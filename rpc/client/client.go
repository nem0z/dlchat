package client

import "fmt"

func Process(args []string) {
	req, err := CreateReq(args)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	resp, err := Send("http://localhost:9999", req)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	fmt.Println()
	resp.Disp()
	fmt.Println()
}
