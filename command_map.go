package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"errors"
)

type Location struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

func commandGetMap(cfg *Config) error {
	url := cfg.Next
	if url == "" {
		return errors.New("No next URL available")
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	location := Location{}
	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return err
	}

	cfg.Next = location.Next
	cfg.Previous = location.Previous
	for _, result := range location.Results {
		fmt.Println(result.Name)
	}
	return nil
	
}

func commandGetMapBack(cfg *Config) error {
	url := cfg.Previous
	if url == "" {
		return errors.New("No previous URL available")
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	location := Location{}
	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return err
	}

	cfg.Next = location.Next
	cfg.Previous = location.Previous
	for _, result := range location.Results {
		fmt.Println(result.Name)
	}
	return nil
	
}