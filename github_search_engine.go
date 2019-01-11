package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const usersURL = "https://api.github.com/search/users"

type Github interface {
	Search(string)
}

type GithubSearchEngine struct {
	Users []*User `json:"items"`
}

func (result *GithubSearchEngine) Search(name string) error {
	resp, err := http.Get(usersURL + "?q=" + name)

	if err != nil {
		return fmt.Errorf("error executing query")
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("error running request")
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return fmt.Errorf("error unmarshall response")
	}

	return nil
}
