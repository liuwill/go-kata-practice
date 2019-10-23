package weekly_contest

import (
	"sort"
	"testing"
)

func compareStringListOrder(source []string, target []string) bool {
	if len(source) != len(target) {
		return false
	}

	sort.Strings(source)
	sort.Strings(target)

	for index := 0; index < len(source); index++ {
		if source[index] != target[index] {
			return false
		}
	}

	return true
}

func Test_RemoveSubfolders(t *testing.T) {
	folderList := [][]string{
		{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"},
		{"/a", "/a/b/c", "/a/b/d"},
		{"/a/b/c", "/a/b/ca", "/a/b/d"},
	}
	output := [][]string{
		{"/a", "/c/d", "/c/f"},
		{"/a"},
		{"/a/b/c", "/a/b/ca", "/a/b/d"},
	}

	for i, folders := range folderList {
		expect := output[i]

		result := removeSubfoldersSlow(folders)
		if !compareStringListOrder(expect, result) {
			t.Error("Run Test_RemoveSubfoldersSlow Fail", i, result)
		}

		target := removeSubfolders(folders)
		if !compareStringListOrder(expect, target) {
			t.Error("Run Test_RemoveSubfolders Fail", i, target)
		}

		t.Log("Run Test_RemoveSubfolders Success")
	}
}
