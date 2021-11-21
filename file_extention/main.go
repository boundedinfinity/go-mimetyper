package file_extention

import "github.com/boundedinfinity/mimetyper/mime_type"

func MimeType(v FileExtention) mime_type.MimeType {
	s := m[v.String()]
	mt, _ := mime_type.Parse(s)
	return mt
}

func DetectMimeType(n string) (mime_type.MimeType, error) {
	var z mime_type.MimeType
	ext, err := Parse(n)

	if err != nil {
		return z, err
	}

	return MimeType(ext), nil
}
