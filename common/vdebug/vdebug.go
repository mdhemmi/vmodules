package vdebug

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
)

func Debug(data string, kind string, task string) {
	switch task {
	case "start":
		fmt.Println("")
		fmt.Println("DEBUG OUTPUT START")
	case "print":
		if kind == "json" {
			var obj map[string]interface{}
			json.Unmarshal([]byte(data), &obj)
			f := colorjson.NewFormatter()
			f.Indent = 4
			s, _ := f.Marshal(obj)
			//s, _ := colorjson.Marshal(obj)
			fmt.Println(string(s))
		} else {
			fmt.Println(data)
		}
	case "end":
		fmt.Println("")
		fmt.Println("DEBUG OUTPUT END")
		fmt.Println("")
	}
}
