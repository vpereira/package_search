package main

type Record struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Release  string `json:"release"`
	Location string `json:"location"`
	Summary  string `json:"summary"`
}
