package file_extention

import (
	"fmt"
	"path/filepath"

	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

func FromExt(ext string) (mime_type.MimeType, error) {
	var typedExt FileExtention
	var err error

	typedExt, err = FileExtentions.Parse(ext)

	if err != nil {
		return mime_type.MimeTypes.Unkown, err
	}

	if mt, ok := ext2mt[typedExt]; ok {
		return mt, nil
	} else {
		return mime_type.MimeTypes.Unkown, fmt.Errorf("can't map extention %v to a MIME type", ext)
	}
}

func FromPath(path string) (mime_type.MimeType, error) {
	ext := filepath.Ext(path)
	return FromExt(ext)
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
