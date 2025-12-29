package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"sort"
	"strconv"
	"strings"
)

// HashSet is a generic set based on a hash table (map).
type HashSet[T comparable] struct {
	m map[T]struct{}
}

// New creates a new HashSet.
func New[T comparable]() *HashSet[T] {
	return &HashSet[T]{m: make(map[T]struct{})}
}

// InitWith creates a new HashSet initialized with vals.
func InitWith[T comparable](vals ...T) *HashSet[T] {
	hs := New[T]()
	for _, v := range vals {
		hs.Add(v)
	}
	return hs
}

// Add adds a value to the set.
func (hs *HashSet[T]) Add(val T) {
	hs.m[val] = struct{}{}
}

// Contains reports whether the set contains the given value.
func (hs *HashSet[T]) Contains(val T) bool {
	_, ok := hs.m[val]
	return ok
}

// Len returns the size/length of the set - the number of values it contains.
func (hs *HashSet[T]) Len() int {
	return len(hs.m)
}

// Delete removes a value from the set; if the value doesn't exist in the
// set, this is a no-op.
func (hs *HashSet[T]) Delete(val T) {
	delete(hs.m, val)
}

// All returns an iterator over all the values in the set.
func (hs *HashSet[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range hs.m {
			if !yield(v) {
				return
			}
		}
	}
}

// Union returns the set union of hs with other. It creates a new set.
func (hs *HashSet[T]) Union(other *HashSet[T]) *HashSet[T] {
	result := New[T]()
	for v := range hs.m {
		result.Add(v)
	}
	for v := range other.m {
		result.Add(v)
	}
	return result
}

// Intersection returns the set intersection of hs with other. It creates a
// new set.
func (hs *HashSet[T]) Intersection(other *HashSet[T]) *HashSet[T] {
	result := New[T]()
	for v := range hs.m {
		if other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

// Difference returns the set difference hs - other. It creates a new set.
func (hs *HashSet[T]) Difference(other *HashSet[T]) *HashSet[T] {
	result := New[T]()
	for v := range hs.m {
		if !other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func parse_products_and_rules(lines []string) ([]int, [][2]int) {
	var products []int
	var rules [][2]int

	for _, line := range lines {
		if strings.Contains(line, "-") {
			rule := strings.Split(line, "-")
			rule_1, _ := strconv.Atoi(strings.Trim(rule[0], " "))
			rule_2, _ := strconv.Atoi(strings.Trim(rule[1], " "))
			rules = append(rules, [2]int{rule_1, rule_2})
		} else if !strings.Contains(line, "-") && line != "" {
			product, _ := strconv.Atoi(strings.Trim(line, " "))
			products = append(products, product)
		}
	}
	return products, rules
}

func open_file(fname string) ([]string, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func part1(products []int, rules [][2]int) int {
	isFreshCount := 0
	for _, product := range products {

		for _, rule := range rules {

			if product >= rule[0] && product <= rule[1] {
				isFreshCount++
				break
			}
		}

	}

	return isFreshCount
}

func part2(rule [][2]int) int {
	var matchingRanges [][2]int
	sort.Slice(rule, func(i, j int) bool {
		return rule[i][0] < rule[j][0]
	})

	for {
		if len(rule) == 1 {
			matchingRanges = append(matchingRanges, rule...)
			break
		}

		// pop off both to compare
		first := rule[0]
		second := rule[1]
		rule = rule[2:]

		// ranges do not intersect
		if second[0] > first[1] {

			matchingRanges = append(matchingRanges, first)
			// add the second item back into the front of the queue
			rule = append([][2]int{{second[0], second[1]}}, rule...)
		} else {
			// ranges intersect redo
			rule = append([][2]int{{first[0], max(first[1], second[1])}}, rule...)
		}
	}

	//count total
	sum := 0
	for _, r := range matchingRanges {
		sum += (r[1] - r[0]) + 1
	}
	return sum
}
func main() {

	workCases, err := open_file("main.txt")

	if err == nil {
		products, rules := parse_products_and_rules(workCases)
		part1 := part1(products, rules)

		fmt.Printf("Part 1: %d\n", part1)

		part2Set := part2(rules)

		fmt.Printf("Part 2: %d\n", part2Set)

	}

}
