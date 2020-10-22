package api

type Options struct {
	log logger
}

type Option func(*Options)

func WithLog(l logger) Option {
	return func(args *Options) {
		args.log = l
	}
}
