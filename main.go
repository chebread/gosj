package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// TODO: json 마샬링 해보기.

func main() {
	x := struct {
		Name  string `json:"released"` // export 되어야만 마샬링이 가능합니다.
		IsMan bool   `json:"age,omitempty"`
	}{
		Name:  "Alice",
		IsMan: false,
	}
	data, err := json.Marshal(x)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
