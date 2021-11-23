package file_extention

import (
	"errors"
	"strings"

	"github.com/boundedinfinity/mimetyper/mime_type"
)

func MimeType(v FileExtention) mime_type.MimeType {
	s := m[v.String()]
	mt, _ := mime_type.Parse(s)
	return mt
}

func DetectMimeType(n string) (mime_type.MimeType, error) {
	var ext FileExtention
	var err error

	ext, err = Parse(n)

	if err != nil && errors.Is(err, ErrFileExtentionInvalid) {
		for _, xext := range All {
			if strings.HasSuffix(n, xext.String()) {
				ext = xext
				err = nil
				break
			}
		}
	}

	if err != nil {
		return mime_type.Unkown, err
	}

	return MimeType(ext), nil
}
