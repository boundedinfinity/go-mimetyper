package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/boundedinfinity/optional"
	"gopkg.in/yaml.v3"
)

type data struct {
	Description   optional.StringOptional `json:"description" yaml:"description"`
	MimeType      string                  `json:"mime-type" yaml:"mime-type"`
	MimeTypeAlt   []string                `json:"mime-type-alt" yaml:"mime-type-alt"`
	FileExtention []string                `json:"file-extention" yaml:"file-extention"`
}

type byMimeType []data

func (t byMimeType) Len() int      { return len(t) }
func (t byMimeType) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t byMimeType) Less(i, j int) bool {
	return t[i].MimeType < t[j].MimeType
}

const (
	dataYaml           = "../data.yaml"
	mimeTypeText       = "../mime-types.txt"
	extentionsText     = "../file-extentions.txt"
	mimeTypeMapGo      = "../mime_type/map.go"
	fileExtentionMapGo = "../file_extention/map.go"
	indent             = 4
)

func main() {
	var data []data

	if err := readDataYaml(&data); err != nil {
		handleError(err)
	}

	var mimeTypes []string

	if err := readMimeTypesText(&mimeTypes); err != nil {
		handleError(err)
	}

	var extentions []string

	if err := readExtentionsText(&extentions); err != nil {
		handleError(err)
	}

	if err := process1(&data, &mimeTypes, &extentions); err != nil {
		handleError(err)
	}

	if err := writeDataYaml(data); err != nil {
		handleError(err)
	}

	if err := writeMimeTypesText(mimeTypes); err != nil {
		handleError(err)
	}

	if err := writeExtentionsText(extentions); err != nil {
		handleError(err)
	}

	if err := process2(data); err != nil {
		handleError(err)
	}
}

func process2(data []data) error {
	mimeType2MimeType := `
    package mime_type

    var (
        m  = map[string]string {
            {{ range $i := . }}
                "{{ $i.MimeType }}": "{{ $i.MimeType }}",
                {{- range $j := $i.MimeTypeAlt }}
                "{{ $j }}": "{{ $i.MimeType }}",
                {{- end }}
            {{- end }}
        }
    )
    `

	if err := writeTemplate(mimeTypeMapGo, mimeType2MimeType, data); err != nil {
		return err
	}

	fileExtention2MimeType := `
    package file_extention
    
    var (
        m  = map[string]string {
            {{ range $i := . }}
            {{- range $j := $i.FileExtention }}
            "{{ $j }}": "{{ $i.MimeType }}",
            {{- end }}
            {{- end }}
        }
    )
    `

	if err := writeTemplate(fileExtentionMapGo, fileExtention2MimeType, data); err != nil {
		return err
	}

	return nil
}

func writeTemplate(p, t string, data []data) error {
	tmpl, err := template.New("").Parse(t)

	if err != nil {
		return err
	}

	var bs bytes.Buffer

	if err := tmpl.Execute(&bs, data); err != nil {
		return err
	}

	bs2, err := format.Source(bs.Bytes())

	if err != nil {
		return err
	}

	file, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		return err
	}

	defer file.Close()

	out := bufio.NewWriter(file)
	defer out.Flush()

	out.WriteString(string(bs2))

	return nil
}

func process1(data *[]data, mimeTypes *[]string, extentions *[]string) error {
	for _, d := range *data {
		sort.Strings(d.MimeTypeAlt)
		sort.Strings(d.FileExtention)
	}

	sort.Sort(byMimeType(*data))

	for _, d := range *data {
		optAppend(mimeTypes, optional.NewStringValue(d.MimeType))

		for _, v := range d.MimeTypeAlt {
			optAppend(mimeTypes, optional.NewStringValue(v))
		}

		for _, v := range d.FileExtention {
			optAppend(extentions, optional.NewStringValue(v))
		}
	}

	sort.Strings(*mimeTypes)
	sort.Strings(*extentions)

	return nil
}

func optAppend(vs *[]string, o optional.StringOptional) {
	if o.IsEmpty() {
		return
	}

	contains := false

	for _, v := range *vs {
		if v == o.Get() {
			contains = true
			fmt.Printf("duplicate: %v\n", o.Get())
			break
		}
	}

	if !contains {
		*vs = append(*vs, o.Get())
	}
}

func readMimeTypesText(vs *[]string) error {
	return readLinesFromFile(mimeTypeText, vs)
}

func readExtentionsText(vs *[]string) error {
	return readLinesFromFile(extentionsText, vs)
}

func readLinesFromFile(p string, vs *[]string) error {
	file, err := os.Open(p)

	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return nil
		}

		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		*vs = append(*vs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func readDataYaml(vs *[]data) error {
	bs, err := ioutil.ReadFile(dataYaml)

	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(bs, &vs); err != nil {
		return err
	}

	return nil
}

func writeDataYaml(vs []data) error {
	var bs bytes.Buffer
	enc := yaml.NewEncoder(&bs)
	enc.SetIndent(indent)

	if err := enc.Encode(vs); err != nil {
		return err
	}

	if err := ioutil.WriteFile(dataYaml, bs.Bytes(), 0755); err != nil {
		return err
	}

	return nil
}

func writeMimeTypesText(vs []string) error {
	return writeLinesToFile(mimeTypeText, vs)
}

func writeExtentionsText(vs []string) error {
	return writeLinesToFile(extentionsText, vs)
}

func writeLinesToFile(p string, vs []string) error {
	file, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		return err
	}

	defer file.Close()

	out := bufio.NewWriter(file)
	defer out.Flush()

	for _, v := range vs {
		out.WriteString(fmt.Sprintf("%v\n", v))
	}

	return nil
}

func handleError(err error) {
	fmt.Print(err.Error())
	os.Exit(1)
}
