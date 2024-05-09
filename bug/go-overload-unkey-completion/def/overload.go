package def

const GopPackage = true

type N struct {
}

func (m *N) OnKey__0(a string, fn func()) {
}

func (m *N) OnKey__1(a string, fn func(key string)) {
}

func (m *N) OnKey__2(a []string, fn func()) {
}

func (m *N) OnKey__3(a []string, fn func(key string)) {
}

func _() {
	// n := &N{}
	// n.OnKey__0()
	// N.OnKey__0()
}
