package main

import (
	"encoding/json"
	"os"
)

type Settings struct {
	Name     string        `json:"name"`
	NotFound NotFoundModel `json:"notFound"`
	Port     int           `json:"port"`
	Path     []PathModel   `json:"paths"`
}

type NotFoundModel struct {
	Status int                    `json:"status"`
	Body   map[string]interface{} `json:"body"`
}

type PathModel struct {
	Path    string        `json:"path"`
	Methods []MethodModel `json:"methods"`
}

type MethodModel struct {
	Method   string        `json:"method"`
	Response ResponseModel `json:"response"`
}

type ResponseModel struct {
	Status  int                    `json:"status"`
	Body    map[string]interface{} `json:"body"`
	Headers map[string]string      `json:"headers"`
}

func GetSettings(path string) ([]Settings, error) {
	file, err := os.Open(path)
	if err != nil {
		return []Settings{}, err
	}
	defer file.Close()

	var settings []Settings
	err = json.NewDecoder(file).Decode(&settings)
	if err != nil {
		return []Settings{}, err
	}

	return settings, nil
}
