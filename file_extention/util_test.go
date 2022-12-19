package file_extention_test

import (
	"testing"

	"github.com/boundedinfinity/go-mimetyper/file_extention"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/stretchr/testify/assert"
)

func Test_Extention2MimeType_valid(t *testing.T) {
	input := file_extention.Aab.String()
	expected := mime_type.ApplicationXAuthorwareBin
	actual, err := file_extention.FromExt(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Extention2MimeType_invalid(t *testing.T) {
	input := ".turd"
	_, err := file_extention.FromExt(input)

	assert.NotNil(t, err)
}

func Test_MimeType2Extention_valid(t *testing.T) {
	expected := []file_extention.FileExtention{file_extention.Yaml, file_extention.Yml}
	actual, err := file_extention.GetExts(mime_type.ApplicationYaml)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
