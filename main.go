package main

import (
	"os"
)

func main() {
	// Open the device file for writing
	file, err := os.OpenFile("/dev/one", os.O_WRONLY, 0666)
	if err != nil {
		println("Error:", err)
		return
	}
	defer file.Close()

	// 0x01 to the device file
	for {
		_, err := file.Write([]byte{0x01})
		if err != nil {
			// If there's an error writing to the device file, print the error and exit
			println("Error:", err)
			return
		}
	}
}
