package state

import (
	"io/ioutil"
	"os"
)

func InitState() error {
	err := ioutil.WriteFile("state", []byte("nada"), 0)
	if err != nil {
		return err
	}
	return nil
}

func ReadState() (string, error) {
	state, err := ioutil.ReadFile("state")
	return string(state), err
}

func WriteState(state string) error {
	err := os.Remove("state")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("state", []byte(state), 0)
	if err != nil {
		return err
	}
	return nil
}
