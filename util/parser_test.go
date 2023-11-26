package util

import (
	"strings"
	"testing"
)

var input = `> hello`

func TestParse(t *testing.T) {
	var (
		b strings.Builder
		msg_read bool
		not_fst_space bool
	)
	for _, v := range input {
		if msg_read {
			if !not_fst_space && v == ' ' || v == '	' {
				not_fst_space = true
			} else {
				not_fst_space = true
				b.WriteRune(v)
			}
			
			if v == '\n' || v == '.' {
				msg_read = false
				not_fst_space = false
			}
			continue;
		}
		switch v {
		case '>':
			msg_read = true
		}
	}

	t.Log(b.String())
}
