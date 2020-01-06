package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// fmt.Println("Response: ", resp)

	//Make a byte slice which will be passed to Read function under Reader Interface
	//We declare the size of the slice, though we know that slice is dynamic and can grow and shrink as required
	//However, Read function is not designed which can increase or decrease the byte slice accordingly
	//instead it will read and shove that much amount of data which can accommodate in slice.
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println("ResponseBody:->", string(bs))
	//We can simplify the above three lines of code, where we are fetching the response
	io.Copy(os.Stdout, resp.Body)
}
