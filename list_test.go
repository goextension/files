package files

import "testing"

func TestList(t *testing.T) {
	files, e := List("d:\\", "mp4", 1)
	if e != nil {
		t.Fatal(e)
	}
	t.Log(files)
}
