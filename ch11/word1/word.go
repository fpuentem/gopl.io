// O pacote word ofrece utilitarios para jogos de palavra.
package word

// IsPalindrome informa se s e lida do mesmo jeito de frente para tras e
// de tras para a frente
// (Nossa primeira tentativa.)
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i]{
			return false
		}
	}
	return true
}
