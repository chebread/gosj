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
	data, err := json.Marshal(x) // byte slice
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	var y struct{ Released string }
	// 필드 태그로 인해 released로 바뀌었으므로,
	// 받을 필드명을 released로 해야 하며,
	// 대문자로 시작해야 하니 Released로 해야 한다.
	// 이렇게 하기 실핟면 필드 태그를 활용하라.
	if err := json.Unmarshal(data, &y); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	fmt.Printf("%s\n", y)
}
