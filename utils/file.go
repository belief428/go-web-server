package utils

import (
	"fmt"
	"os"
)

func Open(name string) (f *os.File, err error) {
	_, err = os.Stat(name)

	if os.IsNotExist(err) {
		return os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	}
	if err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("file %s already exists", name)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsDir(path string) bool {
	s, err := os.Stat(path)

	if err != nil {
		return false
	}
	return s.IsDir()
}

func Mkdir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

func MkdirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func PrepareOutput(path string) error {
	fi, err := os.Stat(path)

	if err != nil {
		if os.IsExist(err) && !fi.IsDir() {
			return err
		}
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, 0777)
		}
	}
	return err
}
