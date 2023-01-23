package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type userModel struct {
	Username string
	Password string
}
type Users []userModel

func (u *Users) add(info userModel) {
	*u = append(*u, info)
	fmt.Println("Added successfully ✅")
}

func (u *Users) Delete(index int) error {
	ls := *u
	if index <= 0 || index > len(ls) {
		return errors.New("Invalid index")
	}

	*u = append(ls[:index-1], ls[index:]...)
	fmt.Println("Delete successful")
	return nil
}

func (u *Users) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, u)
	if err != nil {
		return err
	}
	return nil
}

func (u *Users) Store(filename string) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}
