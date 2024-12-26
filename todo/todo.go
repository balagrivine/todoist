package todo

import (
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Text string
}

func SaveItems(filename string ,items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	
	//fmt.Println(string(b))
	if err = ioutil.WriteFile(filename, b, 0644); err != nil {
		return err
	}
	return nil
}
