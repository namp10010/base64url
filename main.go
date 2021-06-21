package main

import (
	"bufio"
	b64 "encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {
	decode := flag.Bool("d", false, "decode")
	url := flag.Bool("u", false, "url")
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// fmt.Printf("%v, %v, %s \n", *decode, *url, input)

	if *url == false {
		if *decode == false {
			fmt.Println(b64.StdEncoding.EncodeToString([]byte(input)))
		} else {
			sDec, err := b64.StdEncoding.DecodeString(input)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			fmt.Println(string(sDec))
		}
	} else {
		if *decode == false {
			fmt.Println(b64.URLEncoding.EncodeToString([]byte(input)))
		} else {
			uDec, err := b64.URLEncoding.DecodeString(input)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			fmt.Println(string(uDec))
		}
	}
}
