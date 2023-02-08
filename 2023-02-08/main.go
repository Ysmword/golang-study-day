package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Stu struct {
	Name string
}

func main() {
	t := struct {
		time.Time
		N int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}

	m, _ := json.Marshal(t)
	fmt.Printf("%s\n", m)

	l := struct {
		Stu
		N int
	}{
		Stu{Name: "test"},
		5,
	}
	m, _ = json.Marshal(l)
	fmt.Printf("%s\n", m)
}
