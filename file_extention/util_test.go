package file_extention_test

import (
	"testing"

	"github.com/boundedinfinity/go-mimetyper/file_extention"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/stretchr/testify/assert"
)

func Test_Extention2MimeType(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected mime_type.MimeType
		err      error
	}{
		{
			name:     "case 1",
			input:    file_extention.FileExtentions.Aab.String(),
			expected: mime_type.MimeTypes.ApplicationXAuthorwareBin,
			err:      nil,
		},
		{
			name:     "case 2",
			input:    ".turd",
			expected: mime_type.MimeTypes.Unkown,
			err:      file_extention.FileExtentions.Err,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := file_extention.FromExt(tc.input)

			assert.ErrorIs(tt, err, tc.err)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

// func Test_MimeType2Extention_valid(t *testing.T) {
// 	expected := []file_extention.FileExtention{
// 		file_extention.FileExtentions.Yaml,
// 		file_extention.FileExtentions.Yml,
// 	}
// 	actual, err := file_extention.GetExts(mime_type.MimeTypes.ApplicationYaml)

// 	assert.Nil(t, err)
// 	assert.Equal(t, expected, actual)
// }
