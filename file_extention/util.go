package file_extention

import (
	"fmt"

	"github.com/boundedinfinity/mimetyper/mime_type"
)

func Extention2MimeType(n string) (mime_type.MimeType, error) {
	var ext FileExtention
	var err error

	ext, err = Parse(n)

	if err != nil {
		return mime_type.Unkown, err
	}

	if mt, ok := m[ext]; ok {
		return mt, nil
	} else {
		return mime_type.Unkown, fmt.Errorf("can't map extention %v to a MIME type", n)
	}
}
