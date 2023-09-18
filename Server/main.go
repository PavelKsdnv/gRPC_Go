package main

import "google.golang.org/grpc"

func main() {
	v := grpc.Version
	print(v)
}
