package trietree

import "testing"

func TestTrieTree(t *testing.T) {
	tree := NewTrieTree()
	var addTestCases = []struct {
		path  []string
		value string
	}{
		{[]string{"a", "b"}, "ab"},
		{[]string{"a", "b", "c"}, "abc"},
		{[]string{"a", "*"}, "a*"},
		{[]string{"*"}, "*"},
		{[]string{"b", "a"}, "ba"},
	}

	for _, tt := range addTestCases {
		tree.Add(tt.path, tt.value)
	}

	var getTestCases = []struct {
		path     []string
		expected string
	}{
		{[]string{"a", "b"}, "ab"},
		{[]string{"a", "b", "c"}, "abc"},
		{[]string{"a", "b", "c", "d"}, "a*"},
		{[]string{"b", "a"}, "ba"},
		{[]string{"c"}, "*"},
	}

	for _, tt := range getTestCases {
		actual := tree.Get(tt.path)
		if actual != tt.expected {
			t.Errorf("Get(%v) = %v; expected %v", tt.path, actual, tt.expected)
		}
	}

	var delTestCases = []struct {
		path     []string
		expected string
	}{
		{[]string{"a", "b"}, "ab"},
		{[]string{"a", "b", "c"}, "abc"},
		{[]string{"a", "*"}, "a*"},
		{[]string{"*"}, "*"},
		{[]string{"b", "a"}, "ba"},
	}

	for _, tt := range delTestCases {
		actual := tree.Del(tt.path)
		if actual != tt.expected {
			t.Errorf("Del(%v) = %v; expected %v", tt.path, actual, tt.expected)
		}
	}
}
