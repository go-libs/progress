package progress

type HandlerFunc func(current, total, expected int64)

var DefaultHandle = HandlerFunc(func(c, t, e int64) {})

func New() (p *Progress) {
	p = &Progress{Progress: DefaultHandle}
	return
}

type Progress struct {
	Current     int64
	Total       int64
	Expected    int64
	Progress    HandlerFunc
	Finished    bool
	IgnoreTotal bool
}

func (p *Progress) Read(b []byte) (n int, err error) {
	return p.handle(b)
}

func (p *Progress) Write(b []byte) (n int, err error) {
	return p.handle(b)
}

func (p *Progress) handle(b []byte) (n int, err error) {
	n = len(b)
	if p.Finished || n == 0 {
		return
	}
	p.calculate(int64(n))
	p.Progress(p.Current, p.Total, p.Expected)
	return
}

func (p *Progress) calculate(n int64) {
	p.Current += n
	p.Expected = p.Total - p.Current
	if !p.IgnoreTotal && p.Expected < 0 {
		p.Current = p.Total
		p.Expected = 0
		p.Finished = true
	}
}
