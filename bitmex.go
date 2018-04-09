/*
   bitmex.go
       Wrapper for Bitmex exchange API

   Author:
       Pat DePippo <pat@dcrypt.io>
*/
package bitmex

const (
	ApiUrl = "https://www.bitmex.com/api/v1"
)

type Bitmex struct {
	client *Client
}

func New(key, secret string) *Bitmex {
	client := NewClient(key, secret)
	return &Bitmex{client}
}
