// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "firstgoa": hello Resource Client
//
// Command:
// $ goagen
// --design=github.com/phxuecan/FirstGo/firstgoa/design
// --out=$(GOPATH)/src/github.com/phxuecan/FirstGo/firstgoa
// --version=v1.2.0-dirty

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// AnswerHelloPath computes a request path to the answer action of hello.
func AnswerHelloPath() string {

	return fmt.Sprintf("/firstgoa/answer")
}

// answer words and show in response
func (c *Client) AnswerHello(ctx context.Context, path string, payload *Message) (*http.Response, error) {
	req, err := c.NewAnswerHelloRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAnswerHelloRequest create the request corresponding to the answer action endpoint of the hello resource.
func (c *Client) NewAnswerHelloRequest(ctx context.Context, path string, payload *Message) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}

// SayHelloPath computes a request path to the say action of hello.
func SayHelloPath(words string) string {
	param0 := words

	return fmt.Sprintf("/firstgoa/say/%s", param0)
}

// say words and show in response
func (c *Client) SayHello(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSayHelloRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSayHelloRequest create the request corresponding to the say action endpoint of the hello resource.
func (c *Client) NewSayHelloRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
