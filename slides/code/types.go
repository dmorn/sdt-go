package main

import "fmt"

func main() {
	var lang string
	lang = "Go"

	version := "go1.17.2 darwin/amd64"

	lookupTable := map[string]string{
		lang: version,
		"Elixir": "1.12.3 (compiled with Erlang/OTP 24)",
	}

	all := make([]string, 0, len(lookupTable))
	for k, _ := range lookupTable {
		all = append(all, k)
	}

	fmt.Println(all)
}
