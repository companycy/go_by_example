
func newFoo(..., cfg fooConfig) *foo {
	if cfg.Output == nil {
		cfg.Output = ioutil.Discard
	}
}

func main() {
	// don't
	cfg := fooConfig{}
	cfg.Bar = bar
	cfg.Period = 100 * time.Millisecond
	cfg.Output = nil

	// do this, avoid invalid intermediate state
	cfg := fooConfig{
		Bar: bar,
		Period: 100 * time.Millisecond,
		Output: nil,
	}

	// or
	cfg := newFoo(cfg)


	// 
	var (
		payload = flag.String("abc", "efg")
		delay = flag.Duration("delay", "")
	)
	flag.Parse()
}

// info and debug log level are enough ?
// log: go-kit/log
// go install better than go build



