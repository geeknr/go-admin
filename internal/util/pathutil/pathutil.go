// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pathutil

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Clean cleans up given path and returns a relative path that goes straight down.
func Clean(p string) string {
	return strings.Trim(path.Clean("/"+p), "/")
}

func WorkDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(fmt.Sprintf("Read workdir fail: %s", err.Error()))
	}
	return strings.Replace(dir, "\\", "/", -1)
}
