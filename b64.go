package main

import (
	"fmt"
	"os"
	"strconv"
)

var l string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%.8b", binString, c)
	}
	return binString
}
func pad(str string) string {
	ex := (len(str) + (len(str) % 6)) % 6
	for i := 0; i < ex; i++ {
		str += "0"
	}

	return str
}
func removePad(str string) string {
	ex := len(str) % 8
	return str[0 : len(str)-ex]
}
func enc(msg string) string {
	b := stringToBin(msg)

	b = pad(b)
	en := ""
	for i := 0; i < len(b); i += 6 {
		n, err := strconv.ParseInt(b[i:i+6], 2, 64)
		if err != nil {
			fmt.Print(err.Error())
		} else {
			en = fmt.Sprintf("%s%c", en, l[n])
		}

	}

	return en
}
func dec(cip string) string {

	b := ""
	for i := 0; i < len(cip); i++ {
		for j := 0; j < len(l); j++ {
			if l[j] == cip[i] {
				b = fmt.Sprintf("%s%.6b", b, j)

			}
		}

	}

	b = removePad(b)
	dec := ""
	for i := 0; i < len(b); i += 8 {
		n, err := strconv.ParseInt(b[i:i+8], 2, 64)
		if err != nil {
			fmt.Print(err.Error())
		} else {
			dec = fmt.Sprintf("%s%s", dec, string(n))
		}

	}
	return dec
}
func main() {
	if len(os.Args) != 3 {
		fmt.Printf("[-]Usage: %s -d or -e [Message]", os.Args[0])
		return
	} else if os.Args[1] == "-d" {
		cip := os.Args[2]
		fmt.Printf("Cipher: %s\nDecode: %s", cip, dec(cip))
	} else if os.Args[1] == "-e" {
		msg := os.Args[2]
		en := enc(msg)
		fmt.Printf("Message: %s\nEncode: %s", msg, en)
	} else {
		fmt.Printf("[-]Usage: %s -d or -e [Message]", os.Args[0])
		return
	}
}
