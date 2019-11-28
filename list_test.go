package files

import "testing"

func TestList(t *testing.T) {
	files1, e := List("d:\\", "", 0)
	if e != nil {
		t.Fatal(e)
	}
	t.Log(files1)
	files2, e := List("d:\\", ".jpg,.mkv", 2)
	if e != nil {
		t.Fatal(e)
	}
	t.Log(files2)
}
