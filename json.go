package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/TylerBrock/colorjson"
	"io/ioutil"
	"os"
)

func readInput() []byte {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(fmt.Errorf("Error reading info about stdin: %v", err))
	}

	if (info.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("Pretty Print JSON is intended to read from pipe.")
		fmt.Println("Usage: echo '{\"key\": \"value\"}' | ppjson")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(fmt.Errorf("Error reading from stdin: %v", err))
	}

	return b
}

func toJson(b []byte) map[string]interface{} {
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(b), &obj)
	if err != nil {
		panic(fmt.Errorf("Error reading JSON data: %v", err))
	}

	return obj
}

func formatJson(obj map[string]interface{}) string {
	formatter := colorjson.NewFormatter()
	formatter.Indent = 2
	s, err := formatter.Marshal(obj)
	if err != nil {
		panic(fmt.Errorf("error formatting JSON: %v", err))
	}

	return string(s)
}

func main() {
	input := readInput()
	obj := toJson(input)
	s := formatJson(obj)

	fmt.Println(string(s))
}
