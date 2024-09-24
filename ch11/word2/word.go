// O pacote word ofrece utilitarios para jogos de palavra.
package word

import "unicode"

// IsPalindrome informa se s e lida da mesma maneira de frente para
// tras e de tras para a frente.
// A diferevca entre letras maiusculas e minusculas e ignorada, assim
// como os caracteres que nao sao letras.
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r:= range s {
		if unicode.IsLetter(r){
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters{
		if letters[i] != letters[len(letters)-1-i]{
			return false
		}
	}
	return true
}
