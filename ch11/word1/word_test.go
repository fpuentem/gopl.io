package word

import "testing"

func TestPalindrome(t *testing.T){
	if !IsPalindrome("detartrated"){
		t.Error(`IsPalindroeme("detartrated") = false`)
	}
	if !IsPalindrome("kayak"){
		t.Error(`IsPalindroeme("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T){
	if IsPalindrome("palindrome"){
		t.Error(`IsPalindroeme("palindrome") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T){
	if !IsPalindrome("été"){
		t.Error(`IsPalindroeme("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T){
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input){
		t.Errorf(`IsPalindroeme(%q) = false`, input)
	}
}
