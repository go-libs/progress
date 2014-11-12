package progress

type ProgressWriter struct {
	Current, Total, Expected int64
	Progress                 func(current, total, expected int64)
}

func (p *ProgressWriter) Write(b []byte) (int, error) {
	current := len(b)
	if p.Total != -1 {
		p.calculate(int64(current))
		p.Progress(p.Current, p.Total, p.Expected)
	}
	return current, nil
}

func (p *ProgressWriter) calculate(i int64) {
	p.Current += i
	if p.Total >= p.Current {
		p.Expected = p.Total - p.Current
	}
}
