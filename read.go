package main

import (
	"errors"
)

func ReadFile(name string) ([]byte, error){
	input := os.Args
	file, err := os.ReadFile(name)
	if err = nil{
		return nil, errors.New("file Not Found")
	}

	return file, nil
}