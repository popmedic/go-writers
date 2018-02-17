package linewriter

import "testing"

func TestLineWriter(t *testing.T) {
	exp := `[ "this\n" "is\n" "a\n" "test of\n" "the\n" "line writer\n" "I\n" "build testing\n" "it\tout\n" ]`
	w := NewLineWriter(10)
	w.Write([]byte("this\nis\na\ntest of\nthe\nline writer\nI\nbuild testing\nit\tout\n"))
	if w.String() != exp {
		t.Errorf("expected %s got %s", w, exp, w)
	}
	exp = `[ "test of\n" "the\n" "line writer\n" "I\n" "build testing\n" "it\tout\n" "and here\n" "are some new\n" "lines to keep\n" "your busy\n" ]`
	w.Write([]byte("and here\nare some new\nlines to keep\nyour busy\n"))
	if w.String() != exp {
		t.Errorf("expected %s got %s", w, exp, w)
	}
}
