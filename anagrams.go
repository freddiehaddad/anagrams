// Package anagrams defines various functions used for detecting
// anagrams among sets of strings.
package anagrams

import (
	"slices"
	"strings"
	"sync"
)

// token is an intermediate representation used for organizing
// the input strings passed to Group into anagram groups.
type token struct {
	key   string
	value string
}

// createToken returns a token of the string s where key is the
// the string s sorted in ascending order and value is a copy of
// of s.
func createToken(s string) token {
	slice := strings.Split(s, "")
	slices.Sort(slice)
	k := strings.Join(slice, "")
	t := token{k, s}
	return t
}

// createTokens iterates through words creating one thread for each word
// responsible for tokenizing the word. The token is then written to ch.
// Each thread created is preceeded by a call to wg.Add and a call to
// wg.Done when complete. Whe the function returns it also calls wg.Done.
func createTokens(wg *sync.WaitGroup, ch chan token, words []string) {
	defer wg.Done()
	for _, word := range words {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			t := createToken(s)
			ch <- t
		}(word)
	}
}

// readTokens reads tokens received through ch and records the key/value
// in hash.  When the channel is closed, the wg is signaled via a all to Done.
func readTokens(wg *sync.WaitGroup, ch chan token, hash map[string][]string) {
	defer wg.Done()
	for t := range ch {
		hash[t.key] = append(hash[t.key], t.value)
	}
}

// Group organizes all the strings in s into groups of anagrams and
// returns a slice of slices with the results.
func Group(s []string) [][]string {
	anagrams := [][]string{}
	tokenWg := sync.WaitGroup{}
	hashWg := sync.WaitGroup{}
	hash := make(map[string][]string)
	tokens := make(chan token)

	// create anagram tokens
	tokenWg.Add(1)
	go createTokens(&tokenWg, tokens, s)

	// store anagram tokens in hash map
	hashWg.Add(1)
	go readTokens(&hashWg, tokens, hash)

	// wait for worker threads to complete
	tokenWg.Wait()
	close(tokens)
	hashWg.Wait()

	// build the results slice
	for _, v := range hash {
		anagrams = append(anagrams, v)
	}

	return anagrams
}
