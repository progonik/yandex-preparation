# YR04 — Concurrent sitemap crawler

Suggested time: 60 minutes.

## Context

Starting from one URL, crawl pages on the same website and construct a sitemap. A sitemap records every fetched page and the links discovered on that page.

## Provided API

```go
// GetPage downloads one page.
func GetPage(url string) (text string, err error)

// Hostname returns the hostname portion of a URL.
func Hostname(url string) string

// ParsePage extracts links from downloaded page text.
func ParsePage(text string) []string

type PageResult struct {
	Links []string
	Err   error
}

func BuildSitemap(rootURL string, maxConcurrent int) map[string]PageResult
```

Assume the helper functions are safe for concurrent use. `GetPage` may be slow but eventually returns. Assume links are already normalized absolute URLs, so string equality is sufficient for deduplication.

## Part 1 — Sequential crawler

Implement `BuildSitemap` without concurrency first.

Requirements:

- Fetch the root page.
- Parse every successfully downloaded page.
- Follow a discovered link only if its hostname equals the root hostname.
- Fetch each eligible URL at most once.
- Store the links from every successfully parsed page.
- Store the error for every page that could not be downloaded.
- Do not retry failed pages.
- Links to other hostnames may appear in a page's `Links`, but must not be fetched.

### Example

```text
https://example.com/
  -> https://example.com/about
  -> https://other.test/ad

https://example.com/about
  -> https://example.com/
```

Only the first two pages are fetched. The link to `other.test` may be present in the root page result, but it is not added as a fetched page.

## Part 2 — Bounded concurrent crawler

Use `maxConcurrent` to parallelize downloads.

Additional requirements:

- At most `maxConcurrent` calls to `GetPage` may be active at once.
- Deduplicate a URL before it is scheduled, not after its download finishes.
- Protect all shared state from data races.
- Work is dynamic: processing one page may discover more pages.
- Return only after every scheduled page has finished and no page can discover additional work.
- Internal communication must not deadlock when workers discover new URLs.
- Wait for every goroutine before returning.
- Treat `maxConcurrent < 1` as invalid; agree with the interviewer whether to panic or change the function signature to return an error.

## Clarifying questions

1. How should relative URLs be resolved if inputs are not normalized?
2. Are fragments or different query strings distinct URLs?
3. Should failed pages be retried or terminate the whole crawl?
4. Is there a maximum crawl depth or total page count?
5. Must the result map have deterministic iteration or output order?
6. How should redirects to another hostname be handled?

## What the interviewer will test

- Deduplication under concurrency.
- Enforcement of a concurrency limit.
- Completion detection for dynamically generated work.
- Channel ownership, shutdown, and goroutine joining.

## Provenance

Reconstructed from a [detailed first-person Yandex backend interview report](https://habr.com/ru/articles/1006022/). The normalized-URL and error-recording rules were added to make the practice version deterministic.

For a related harder exercise, see [H01 — Bounded concurrent crawler](../../problems/hard/h01-web-crawler/README.md).
