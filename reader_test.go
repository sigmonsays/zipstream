package zipstream

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"testing"
)

func TestReader(t *testing.T) {
	s := []byte(`<poc><firstName>Juan</firstName></poc>`)

	var wbuf bytes.Buffer
	z := zip.NewWriter(&wbuf)
	for i := 0; i < 2; i++ {
		zw, err := z.Create("tmp")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := zw.Write(s); err != nil {
			t.Fatal(err)
		}
	}

	if err := z.Close(); err != nil {
		t.Fatal(err)
	}

	zr := NewReader(&wbuf)
	for i := 0; i < 2; i++ {
		_, err := zr.Next()
		if err != nil {
			t.Fatal("Embeded file missing")
		}
		s2, err := ioutil.ReadAll(zr)
		if err != nil {
			t.Fatal(err)
		}

		if bytes.Compare(s, s2) != 0 {
			t.Fatal("Decompressed data does not match original")
		}
	}
}
