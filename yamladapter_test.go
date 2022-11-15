package yamladapter_test

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"testing"

	yamladapter "github.com/rwx-yxu/caddy-yaml-adapter"
)

func TestFixtures(t *testing.T) {
	f, err := filepath.Glob("testdata/file*")
	if err != nil {
		t.Fatal(err)
	}
	l := len(f)
	if l == 0 {
		t.Fatal("no test files found")
	}

	sort.Strings(f)

	a := yamladapter.Adapter{}

	for i := 0; i < l; i += 2 {
		//Json file
		jsn, err := os.ReadFile(f[i])
		if err != nil {
			t.Error(err)
			continue
		}

		//Yaml file
		ym, err := os.ReadFile(f[i+1])
		if err != nil {
			t.Error(err)
			continue
		}

		out, _, err := a.Adapt(ym, nil)
		if err != nil {
			t.Error(err)
			continue
		}

		if !bytes.Equal(bytes.TrimSpace(jsn), out) {
			t.Errorf("fixture %d failed", i+1)
			t.Logf("expected: %s\nactual: %s", jsn, out)
		}
	}
}
