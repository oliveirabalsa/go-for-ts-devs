package main

import (
	"fmt"
	"io/ioutil"
	// "os"
)

const filePermissions = 0644

func main() {
	// !INFO: Writing to a file
	data := []byte("Hello, Go!")
	err := ioutil.WriteFile("output-go.txt", data, filePermissions)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File written successfully")
	}

	//!INFO: Reading from a file
	fileData, err := ioutil.ReadFile("output-go.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File content:", string(fileData))
	}

	//!INFO: Deleting a file
	// err = os.Remove("output-go.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("File deleted successfully")
	// }
}
