package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Get() ([]User, error) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var u []User
	err = json.Unmarshal(b, &u)

	return u, err
}
