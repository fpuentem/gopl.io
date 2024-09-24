package main

func TestEcho(t *testing.T){
	var tests = []struct{
		newline bool
		sep string
		args []string
		want string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []}
	}
}
