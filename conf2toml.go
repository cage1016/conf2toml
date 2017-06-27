package conf2toml

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Reference
// github.com/asaskevich/govalidator
const (
	Int   string = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	Float string = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
)

var (
	rxInt   = regexp.MustCompile(Int)
	rxFloat = regexp.MustCompile(Float)
)

func replaceSpace(input string) string {
	return strings.Replace(string(input), " ", "", -1)
}

// isInt
func isInt(str string) bool {
	if len(str) == 0 {
		return true
	}
	return rxInt.MatchString(str)
}

// IsFloat
func isFloat(str string) bool {
	return str != "" && rxFloat.MatchString(str)
}

// isBoolean
func isBoolean(str string) bool {
	if _, err := strconv.ParseBool(str); err == nil {
		return true
	}
	return false
}

// handle header
func handleheader(input []byte) []byte {
	i := replaceSpace(string(input))
	i = strings.Replace(i, "-", "", -1)
	return []byte(i)
}

// handle line
func handleline(input []byte) []byte {
	i := string(input)

	buf := strings.SplitAfterN(i, "=", 2)
	r := []string{
		replaceSpace(strings.Replace(buf[0], "=", "", -1)),
		strings.TrimSpace(buf[1]),
	}

	if len(r[1]) == 0 {
		r[1] = fmt.Sprintf("\"\"")
	} else if isInt(r[1]) || isFloat(r[1]) {
		r[1] = r[1]
	} else if isBoolean(r[1]) {
		b, _ := strconv.ParseBool(r[1])
		r[1] = strconv.FormatBool(b)
	} else {
		r[1] = fmt.Sprintf("\"%s\"", r[1])
	}

	x := fmt.Sprintf("%v=%v", r[0], r[1])

	return []byte(x)
}

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

	lines := bytes.Split(buf, []byte{'\n'})

	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			f.WriteString("\n")
			continue
		}

		switch line[0] {
		case '[':
			line = handleheader(line)
		default:
			line = handleline(line)
		}
		f.WriteString(fmt.Sprintf("%s\n", line))
	}
	f.Seek(0, 0)
	return f
}
