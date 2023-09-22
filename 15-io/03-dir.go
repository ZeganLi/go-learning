package main

import "os"

func main() {
	os.Mkdir("./dir", os.ModePerm)
}
