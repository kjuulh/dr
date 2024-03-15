package main

import "shuttle-extensions-template/cmd"

func main() {
	if err := cmd.RootCmd().Execute(); err != nil {
		panic(err)
	}
}
