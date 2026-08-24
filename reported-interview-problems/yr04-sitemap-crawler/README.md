# YR04 - Bounded concurrent sitemap crawler

Source: [detailed first-person Yandex backend interview report, February 2026](https://habr.com/ru/articles/1006022/).

The interviewer supplies helpers equivalent to:

```go
func GetPage(url string) (text string, err error)
func Hostname(url string) string
func ParsePage(text string) []string
```

Starting from one root URL, construct a sitemap represented by a map from each fetched URL to the links found on that page.

## Reconstructed basic requirements

- Begin with the root page.
- Fetch and parse pages reachable through discovered links.
- Follow links only when their hostname matches the root hostname.
- Do not fetch the same URL repeatedly.
- Preserve enough error information to distinguish failed downloads from successfully parsed pages.
- Return the discovered graph as `map[url]links` or an equivalent structured result.

## Reported concurrency extension

Parallelize page fetching with at most `N` active fetches.

- The work set is dynamic: fetching one page can discover more pages.
- Deduplicate URLs before scheduling them.
- Protect shared result and visited state.
- Do not let workers deadlock while both consuming and producing work.
- Finish only when the pending queue is empty and no worker is still processing a page that could discover more work.
- Join every worker before returning.

## Clarifications to ask

1. How should URLs be normalized before deduplication?
2. Should a failed page be retried, recorded, or fail the entire crawl?
3. Are fragments and query strings distinct pages?
4. Is there a maximum depth or page count?
5. Must result links include external URLs even though they are not fetched?

This is closely represented by the existing exercise [H01 - Bounded concurrent crawler](../../problems/hard/h01-web-crawler/README.md).
