# Wordle Solver

A quick and dirty Wordle solver written in Go. Uses `/usr/share/dict/words` as the list of all valid words.

### Inputs

This program takes 3 inputs as flags:

- `hint` (defaults to `"*****"`) - A string containing the characters that are already known. Any unknown characters should be denoted by an asterisk. This **must** be exactly 5 characters.
- `unplaced` (defaults to `""`) - A string containing the characters that are known to be in the word, but are unplaced.
- `rejected` (defaults to `""`) - A string containing the characters that have been rejected from the word.

### Example Usage

```
go run main.go -hint='*inc*' -unplaced='e' -rejected ='tarsovph'
```

### Unicode

I haven't tested it very thoroughly, but this _should_ be Unicode-safe. This was done defensively as I wasn't sure if `/usr/share/dict/words` has any words with accent characters.

### Performance

It's fast enough to be useful as-is so I haven't tried any real optimizations:

```
Running tool: /usr/local/bin/go test -benchmem -run=^$ -bench ^BenchmarkFindCandidates$ pranj.co/wordle

goos: darwin
goarch: amd64
pkg: pranj.co/wordle
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkFindCandidates-8   	      52	  22512896 ns/op	 3249967 B/op	  261041 allocs/op
PASS
ok  	pranj.co/wordle	2.243s
```
