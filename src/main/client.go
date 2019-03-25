package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	proto "proto"
)



func main() {

	service := micro.NewService(
		micro.Name("palindrome.client"),
	)


	service.Init()

    palindrome := proto.NewPalindromeService("palindrome", service.Client())

	rsp, err := palindrome.FindPalindrome(context.TODO(), &proto.PalindromeRequest{Input:"palindrome"})

    if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Palindrome)
}
