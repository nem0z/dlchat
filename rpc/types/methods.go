package types

import "fmt"

func (resp *Response) Disp() {
	if resp.Err != "" {
		fmt.Printf("Error : %v\n", resp.Err)
	} else {
		fmt.Println(resp.Result)
	}
}
