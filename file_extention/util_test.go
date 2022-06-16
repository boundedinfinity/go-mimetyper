package file_extention_test

import (
	"testing"

	"github.com/boundedinfinity/mimetyper/file_extention"
	"github.com/boundedinfinity/mimetyper/mime_type"
	"github.com/stretchr/testify/assert"
)

func Test_Extention2MimeType_valid(t *testing.T) {
	input := file_extention.Aab.String()
	expected := mime_type.ApplicationXAuthorwareBin
	actual, err := file_extention.Extention2MimeType(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Extention2MimeType_invalid(t *testing.T) {
	input := ".turd"
	_, err := file_extention.Extention2MimeType(input)

	assert.NotNil(t, err)
}
