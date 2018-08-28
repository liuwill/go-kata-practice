package weekly_contest_83

import "testing"

func Test_MaskPIIEmail(t *testing.T) {
	source := "LeetCode@LeetCode.com"
	target := maskPII(source)
	expect := "l*****e@leetcode.com"

	if expect != target {
		t.Error("Translate Test_MaskPIIEmail Fail", target)
	}
	t.Log("Translate Test_MaskPIIEmail Success")
}

func Test_MaskPIIShortEmail(t *testing.T) {
	source := "AB@qq.com"
	target := maskPII(source)
	expect := "a*****b@qq.com"

	if expect != target {
		t.Error("Translate Test_MaskPIIShortEmail Fail", target)
	}
	t.Log("Translate Test_MaskPIIShortEmail Success")
}

func Test_MaskPIIPhone(t *testing.T) {
	source := "1(234)567-890"
	target := maskPII(source)
	expect := "***-***-7890"

	if expect != target {
		t.Error("Translate Test_MaskPIIPhone Fail", target)
	}
	t.Log("Translate Test_MaskPIIPhone Success")
}

func Test_MaskPIIChinaPhone(t *testing.T) {
	source := "86-(10)12345678"
	target := maskPII(source)
	expect := "+**-***-***-5678"

	if expect != target {
		t.Error("Translate Test_MaskPIIChinaPhone Fail", target)
	}
	t.Log("Translate Test_MaskPIIChinaPhone Success")
}
