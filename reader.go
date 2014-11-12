package progress

type ProgressReader struct {
	Current, Total, Expected int64
	Progress                 func(current, total, expected int64)
}

func (p *ProgressReader) Read(b []byte) (int, error) {
	current := len(b)
	if p.Total != -1 {
		p.calculate(int64(current))
		p.Progress(p.Current, p.Total, p.Expected)
	}
	return current, nil
}

func (p *ProgressReader) calculate(i int64) {
	p.Current += i
	if p.Total >= p.Current {
		p.Expected = p.Total - p.Current
	}
}
