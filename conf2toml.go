package conf2toml

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
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
	f, _ := ioutil.TempFile("", "tmp-conf")
	defer os.Remove(f.Name())

	w := bufio.NewWriter(f)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if len(strings.TrimSpace(scanner.Text())) == 0 {
			fmt.Fprintln(w, strings.Replace(scanner.Text(), `\n`, "\n", -1))
		} else {
			fmt.Fprintln(w, transform(scanner.Text()))
		}
	}
	w.Flush()
	f.Seek(0, 0)
	return f
}
