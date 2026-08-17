# A06 - Minimum dictionary segmentation

Target: 50 minutes.

Implement:

```go
func MinimumSegmentation(
    text string,
    dictionary []string,
) (words []string, ok bool)
```

Split the entire `text` into dictionary words while using the minimum possible number of words. Return one minimum segmentation and `true`. If segmentation is impossible, return `nil, false`.

Dictionary entries may be reused. Duplicate dictionary entries have no additional meaning. If several minimum segmentations exist, return any one of them.

All strings contain only lowercase English letters.

## Examples

```text
text = "leetcode"
dictionary = ["leet", "code", "leetcode"]
answer = ["leetcode"], true

text = "catsanddog"
dictionary = ["cat", "cats", "and", "sand", "dog"]
answer = ["cats", "and", "dog"], true
another valid answer = ["cat", "sand", "dog"], true

text = "catsandog"
dictionary = ["cats", "dog", "sand", "and", "cat"]
answer = nil, false

text = ""
dictionary = ["a"]
answer = [], true
```

## Constraints

- `0 <= len(text) <= 5_000`
- `0 <= len(dictionary) <= 20_000`
- Dictionary words are non-empty.
- Every dictionary word has length at most `100`.

Target complexity: `O(len(text) * maxWordLength)` expected time after building a lookup set, plus `O(len(text))` dynamic-programming and reconstruction space.

Besides finding the minimum count, preserve enough information to reconstruct the selected words.
