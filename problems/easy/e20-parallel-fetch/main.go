package main

import (
	"context"
	"errors"
	"sync"
)

type Fetcher interface {
	Fetch(context.Context, string) ([]byte, error)
}

type Job struct {
	Content string
	Index   int
}

type Page struct {
	URL  string
	Body []byte
}

var ErrInvalidWorkersCount = errors.New("workers count must be positive")

func FetchAll(
	ctx context.Context,
	fetcher Fetcher,
	urls []string,
	maxConcurrent int,
) ([]Page, error) {
	if maxConcurrent < 1 {
		return nil, ErrInvalidWorkersCount
	}
	l := len(urls)
	if l == 0 {
		return []Page{}, nil
	}

	result := make([]Page, l)
	jobs := make(chan Job)
	var wg sync.WaitGroup
	var finalError error

	maxConcurrent = min(maxConcurrent, l)

	var errOnce sync.Once

	wg.Add(maxConcurrent)

	workerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	for range maxConcurrent {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-workerCtx.Done():
					return
				case job, ok := <-jobs:
					if !ok {
						return
					}
					if err := workerCtx.Err(); err != nil {
						return
					}

					body, err := fetcher.Fetch(workerCtx, job.Content)
					if err != nil {
						errOnce.Do(func() {
							finalError = err
							cancel()
						})
						return
					}
					result[job.Index] = Page{
						URL:  job.Content,
						Body: body,
					}
				}
			}
		}()
	}

broke:
	for i, url := range urls {
		select {
		case <-workerCtx.Done():
			break broke
		case jobs <- Job{
			Content: url,
			Index:   i,
		}:
		}
	}

	close(jobs)
	wg.Wait()

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if finalError != nil {
		return nil, finalError
	}

	return result, nil
}
