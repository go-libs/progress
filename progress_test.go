package progress

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/bmizerany/assert"
)
import "testing"

func TestNew(t *testing.T) {
	p := NewWriter()
	zero := int64(0)
	assert.Equal(t, zero, p.Current)
	assert.Equal(t, zero, p.Total)
	assert.Equal(t, zero, p.Expected)
}

func TestProgress(t *testing.T) {
	filename := "progress_test.go"
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fs, err := os.Stat(filename)
	if err != nil {
		log.Fatalln(err)
	}

	p := NewWriter()
	p.Total = fs.Size()
	p.Progress = func(current, total, expected int64) {
		log.Println(current, total, expected)
		assert.Equal(t, true, current <= total)
	}

	b := new(bytes.Buffer)
	w := io.MultiWriter(p, b)
	_, err = io.Copy(w, f)
	if err != nil {
		log.Fatalln(err)
	}
}
