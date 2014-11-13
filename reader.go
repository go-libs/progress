package progress

import "io"

type ProgressReader struct {
	Current, Total, Expected int64
	Finished                 bool
	Progress                 func(current, total, expected int64)
}

func (p *ProgressReader) Read(b []byte) (n int, err error) {
	if p.Finished {
		return 0, io.EOF
	}
	n = len(b)
	p.Finished = p.calculate(int64(n))
	p.Progress(p.Current, p.Total, p.Expected)
	return
}

func (p *ProgressReader) calculate(n int64) bool {
	p.Current += n
	if p.Current > p.Total {
		p.Current = p.Total
	}
	p.Expected = p.Total - p.Current
	return p.Current == p.Total
}
