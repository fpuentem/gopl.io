//  O pacote memo ofrece uma memoizacao
// segura para concorrencia de uma funcao do tipo Func
package memo

// Um Memo fez cache dos resultados da chamada a uma Func.
type Memo struct{
	f  Func
	cache map[string]result
}

// Func e o tipo da funcao para memoizar
type Func func(key string) (interface{}, error)
type result struct{
	value interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// NOTA: nao e seguro para concorrencia
func (memo *Memo) Get(key string) (interface{}, error){
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
