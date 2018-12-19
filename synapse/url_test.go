package synapse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/********** TESTS **********/

func Test_BuildURL(t *testing.T) {
	basePath := "https://randomsite.com"
	path1 := "test1"
	path2 := "test2"
	path3 := "test3"
	fullURL := basePath + "/" + path1 + "/" + path2 + "/" + path3

	testURL := buildURL(basePath, path1, path2, path3)

	assert.Equal(t, fullURL, testURL)
}
