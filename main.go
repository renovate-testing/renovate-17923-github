package main

import (
	subgithub "github.com/lunarway/go-renovate-submodules/submodule-github"
)

func main() {
	println("Hello from root")

	subgithub.PrintHello()
}
