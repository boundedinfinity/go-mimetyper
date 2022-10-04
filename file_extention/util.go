package file_extention

import (
	"fmt"

	"github.com/boundedinfinity/mimetyper/mime_type"
)

func GetMimeType(ext string) (mime_type.MimeType, error) {
	var typedExt FileExtention
	var err error

	typedExt, err = Parse(ext)

	if err != nil {
		return mime_type.Unkown, err
	}

	if mt, ok := ext2mt[typedExt]; ok {
		return mt, nil
	} else {
		return mime_type.Unkown, fmt.Errorf("can't map extention %v to a MIME type", ext)
	}
}

func GetExts(mimetypes ...mime_type.MimeType) ([]FileExtention, error) {
	var out []FileExtention
	var resolvedtypes []mime_type.MimeType

	for _, mt := range mimetypes {
		realMt := mime_type.ResolveMimeType(mt)
		resolvedtypes = append(resolvedtypes, realMt)
	}

	for _, mt := range resolvedtypes {
		list, ok := mt2ext[mt]

		if !ok {
			return out, fmt.Errorf("no extentions for MIME type %v", mt)
		}

		out = append(out, list...)
	}

	return out, nil
}
