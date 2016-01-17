/*
  Functions for temporary file-related operations

  @author Colorado Reed (colorado.j.reed _at_ gmail.com)
*/

package temp

import (
	"os"
	"path/filepath"
)

// Filename returns the name for a temporary file in
// the system's temporary directory
func Filename() string {
	dir := os.TempDir()
	var filename string
	for i := 0; i < 10000; i++ {
		filename = filepath.Join(dir, randShortString())
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			break
		}
	}
	return filename
}
