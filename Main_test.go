package main

import "testing"

// at 37abb4760815343adbdd2e50090131f2ddcd3e11
//PASS
//1	16794929234 ns/op
//ok  	rasp_parser	16.814s

func BenchmarkMain(b *testing.B)  {
	for n:=0; n < b.N; n++ {
		main()
	}
}