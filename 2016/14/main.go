package main

import (
	"crypto/md5"
	"fmt"
	"strconv"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

var salt = "ahsbgdzn"

type Candidate struct {
	Hash  string
	Digit string
	Index int
}

func (c Candidate) String() string {
	return fmt.Sprintf("%s triple: %s index: %d", c.Hash, c.Digit, c.Index)
}

const rounds = 2017 // 2016 + initial hash

func generateHash(index int) string {
	result := salt + strconv.Itoa(index)
	for i := 0; i < rounds; i++ {
		result = fmt.Sprintf("%x", md5.Sum([]byte(result)))
	}
	return result
}

func main() {
	keys := make([]*Candidate, 0)
	candidates := make(map[string][]*Candidate)

	var (
		letterx3 = pcre.MustCompile(`(.)\1{2}`, 0)
		letterx5 = pcre.MustCompile(`(.)\1{4}`, 0)
	)
	for i := 0; ; i++ {
		hash := generateHash(i)
		matcher := letterx3.MatcherString(hash, 0)
		if !matcher.Matches() {
			continue
		}

		candidate := &Candidate{
			Hash:  hash,
			Digit: matcher.GroupString(1),
			Index: i,
		}

		matcher = letterx5.MatcherString(hash, 0)
		if matcher.Matches() {
			//fmt.Printf("Found 5x in: %s (index %d)\n", hash, i)
			digit := matcher.GroupString(1)
			for _, c := range candidates[digit] {
				//fmt.Println("  Considering", c)
				if i-c.Index <= 1000 {
					keys = append(keys, c)
				}
			}
			delete(candidates, digit)
			//fmt.Println("keys so far:", len(keys))
		}

		// triple and 5x in same hash doesn't count so has to be appended afterwards
		candidates[candidate.Digit] = append(candidates[candidate.Digit], candidate)

		if len(keys) >= 64 {
			fmt.Println("Found 64th key at index", keys[63].Index)
			break
		}
	}
}
