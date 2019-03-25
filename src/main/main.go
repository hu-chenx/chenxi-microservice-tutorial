package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	proto "proto"
)

type Palindrome struct{}

func (g *Palindrome) FindPalindrome(ctx context.Context, req *proto.PalindromeRequest, rsp *proto.PalindromeResponse) error {
	rsp.Palindrome = "Hello " + req.Input
	return nil
}

func main() {

	service := micro.NewService(
		micro.Name("palindrome"),
	)


	service.Init()

	proto.RegisterPalindromeHandler(service.Server(), new(Palindrome))


	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
