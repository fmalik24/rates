package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type IFileSystem interface {
	getDataFromFileSystem() []byte
	saveDataToFileSystem(data []byte)
}

type repo struct{}

func (repo repo) getDataFromFileSystem() []byte {
	return getDataFromFileSystem()
}

func (repo repo) saveDataToFileSystem(data []byte) {
	saveDataToFileSystem(data)
}

func getDataFromFileSystem() []byte {
	jsonFile, err := os.Open("rate.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue

}

func saveDataToFileSystem(ratesJSON []byte) {
	ratesFile, err := os.Create("rate.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = ratesFile.Write(ratesJSON)
	if err != nil {
		fmt.Println(err)
		ratesFile.Close()
		return
	}

	err = ratesFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
