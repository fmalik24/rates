package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type someHack interface {
	getDataFromFileSystem() []byte
}

type repo struct{}

func (repo repo) getDataFromFileSystem() []byte {
	return getDataFromFileSystem()
}

func getDataFromFileSystem() []byte {
	jsonFile, err := os.Open("user.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue

}
