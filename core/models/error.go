package models

type Error struct {
	Message    string
	Extensions map[string]interface{}
}
