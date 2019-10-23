package weekly_contest

import "testing"

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
		target := removeSubfolders(folders)

		if !compareStringList(expect, target) {
			t.Error("Run Test_RemoveSubfolders Fail", i, target)
		}
		t.Log("Run Test_RemoveSubfolders Success")
	}
}
