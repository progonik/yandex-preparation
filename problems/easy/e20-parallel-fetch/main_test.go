package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
)

var ErrSpecial = errors.New("special error")

type HTTPFetcher struct{}

func NewHTTPFetcher() *HTTPFetcher {
	return &HTTPFetcher{}
}

func (f *HTTPFetcher) Fetch(ctx context.Context, s string) ([]byte, error) {
	if strings.Contains(s, "error") {
		return nil, errors.New(s)
	}

	return fmt.Appendf([]byte("byte-body-for-%s"), s), nil
}

func TestCorrectBehavior(t *testing.T) {
	fetcher := NewHTTPFetcher()
	testCases := []struct {
		Name          string
		Fetcher       Fetcher
		Urls          []string
		MaxConcurrent int
	}{
		{
			Name:          "Test-With-Less-Workers-Than-URLs",
			Fetcher:       fetcher,
			Urls:          []string{"test-url-1", "test-url-2", "test-url-3", "test-url-4", "test-url-5", "test-url-6", "test-url-7", "test-url-8"},
			MaxConcurrent: 3,
		},
		{
			Name:          "Test-With-More-Workers-Than-URLs",
			Fetcher:       fetcher,
			Urls:          []string{"test-url-1", "test-url-2", "test-url-3", "test-url-4", "test-url-5", "test-url-6", "test-url-7", "test-url-8"},
			MaxConcurrent: 10,
		},
		{
			Name:          "Test-With-Empty-URLs",
			Fetcher:       fetcher,
			Urls:          []string{},
			MaxConcurrent: 3,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result, err := FetchAll(
				context.Background(),
				testCase.Fetcher,
				testCase.Urls,
				testCase.MaxConcurrent,
			)

			if err != nil {
				t.Fatalf("execution returned an unexpected error %s", err.Error())
			}
			if len(result) != len(testCase.Urls) {
				t.Fatalf("execution returned invalid number of result, expected %d", len(testCase.Urls))
			}
		})
	}
}

func TestErrorCases(t *testing.T) {
	fetcher := NewHTTPFetcher()
	testCases := []struct {
		Name          string
		Fetcher       Fetcher
		Urls          []string
		MaxConcurrent int
	}{
		{
			Name:          "Error-Test-With-Less-Workers-Than-URLs",
			Fetcher:       fetcher,
			Urls:          []string{"test-url-1", "test-url-2", "error-test-url-3", "test-url-4", "test-url-5", "test-url-6", "test-url-7", "test-url-8"},
			MaxConcurrent: 3,
		},
		{
			Name:          "Error-Test-With-More-Workers-Than-URLs",
			Fetcher:       fetcher,
			Urls:          []string{"error-test-url-1", "test-url-2", "test-url-3", "test-url-4", "test-url-5", "test-url-6", "test-url-7", "test-url-8"},
			MaxConcurrent: 10,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result, err := FetchAll(context.Background(), fetcher, testCase.Urls, testCase.MaxConcurrent)
			if err == nil {
				t.Fatal("execution could not handle the error")
			}
			if result != nil {
				t.Fatalf("execution returned an unexpected response %v", result)
			}
		})
	}
}

func TestWithMoreThanOneErrorCase(t *testing.T) {
	fetcher := NewHTTPFetcher()

	t.Run("Test-With-Two-Errors", func(t *testing.T) {
		result, err := FetchAll(
			context.Background(),
			fetcher,
			[]string{"test-url-1", "test-url-2", "test-url-3", "test-error-case-1", "test-url-4", "test-url-5", "test-url-6", "test-with-error-case-2", "test-url-7", "test-url-8"},
			5,
		)

		if err != nil {
			if err.Error() != "test-error-case-1" {
				t.Fatalf("execution returned wrong error: %s", err.Error())
			}
		} else {
			t.Fatal("execution did not return any errors")
		}

		if result != nil {
			t.Fatalf("execution returned an unexpected response %v", result)
		}
	})
}
