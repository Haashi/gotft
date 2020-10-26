package api

import "github.com/haashi/gotft/api/internal"

type Options struct {
	log internal.Logger
	c   internal.HttpClient
}

type Option func(*Options)

func WithLog(l internal.Logger) Option {
	return func(args *Options) {
		args.log = l
	}
}

func WithClient(c internal.HttpClient) Option {
	return func(args *Options) {
		args.c = c
	}
}
