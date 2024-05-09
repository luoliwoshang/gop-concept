package main

import "test.com/def"

func over() {
	n := &def.N{}
	// n.OnKey()
	// 错误：n后面不应该Completion OnKey
	_ = n
}
