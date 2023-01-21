package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type userDetail struct {
	Username string
	Email    string
	Password string
	Note     string
}
type Details []userDetail

func (d *Details) add(detail userDetail) {
	*d = append(*d, detail)
	fmt.Println("Added successfully!")
}

func (d *Details) Delete(index int) error {
	ls := *d
	if index <= 0 || index > len(ls) {
		return errors.New("Invalid index")
	}

	*d = append(ls[:index-1], ls[index:]...)
	fmt.Println("Delete successful")
	return nil
}

func (d *Details) Load(filename string) error {
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
	err = json.Unmarshal(file, d)
	if err != nil {
		return err
	}
	return nil
}

func (d *Details) Store(filename string) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}
