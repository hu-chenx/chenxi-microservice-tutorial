package main

import (
	"context"
	"fmt"
	"strconv"

	micro "github.com/micro/go-micro"
	proto "proto"
)

type Palindrome struct{}

func (g *Palindrome) FindPalindrome(ctx context.Context, req *proto.PalindromeRequest, rsp *proto.PalindromeResponse) error {
	rsp.Palindrome = "Hello " +   strconv.Itoa(longestPalindrome(req.Input))
	return nil
}

func longestPalindrome(input string) int {
    m := make(map[rune]int)
    for _, v := range input {
        m[v]++
    }
    var len int
    var odd bool
    for _, vv := range m {
        len += vv
        if vv % 2 == 1 {
            len--
            odd = true
        }
    }
    if odd {
        len++
    }
    return len
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
