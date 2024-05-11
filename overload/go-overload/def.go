package main

const GopPackage = true

type N struct {
}

func (m *N) OnKey__0(a string, fn func()) {
	fn()
}

func (m *N) OnKey__1(a string, fn func(key string)) {
	fn(a)
}

func (m *N) OnKey__2(a []string, fn func()) {
	fn()
}
