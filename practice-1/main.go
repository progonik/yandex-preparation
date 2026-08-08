package main

import "context"

type Fetcher interface {
	Fetch(ctx context.Context, url string) ([]byte, error)
}

type Page struct {
	URL  string
	Body []byte
}

func FetchAll(
	ctx context.Context,
	fetcher Fetcher,
	urls []string,
	maxConcurrent int,
) ([]Page, error)
