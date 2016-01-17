/*
tempfile_test contains tests for the temp packages temporary file functionality.

@author Colorado Reed (colorado.j.reed _at_ gmail.com)
*/

package temp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Make sure
func TestIfReturnNonExistingFile(t *testing.T) {
	tmpfilename := Filename()
	_, err := os.Stat(tmpfilename)
	assert.True(t, os.IsNotExist(err))
	assert.True(t, len(tmpfilename) > 0, "tmp filename is an empty string")
}
