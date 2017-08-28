package conf2toml

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Normalization(path string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}

	f := NormalizationReader(bytes.NewReader(buf))
	return ioutil.ReadAll(f)
}

func NormalizationReader(input io.Reader) *os.File {
	buf, err := ioutil.ReadAll(input)
	if err != nil {
		return (*os.File)(nil)
	}

	f, _ := ioutil.TempFile("", "tmp-conf")
	defer os.Remove(f.Name())

	if binary.Size(buf) == 0 {
		return f
	}

	lines := bytes.Split(buf, []byte{'\n'})

	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			f.WriteString("\n")
			continue
		}

		f.WriteString(fmt.Sprintf("%s\n", transform(string(line))))
	}
	f.Seek(0, 0)
	return f
}
