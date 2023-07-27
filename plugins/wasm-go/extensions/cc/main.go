package main

import "fmt"

func main() {
	fmt.Println("Hello world.")
}

type CCConfig struct {
	headerRulesMap map[string][]CCRule
	cookieRulesMap map[string][]CCRule
}

type CCRule struct {
}
