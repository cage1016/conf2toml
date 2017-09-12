package conf2toml

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

const (
	StarWithEe string = `^[eE].*`
	Pair       string = `([^=]*)=(.*)`
	Trans      string = "[[:punct:]]|[[:space:]]"
)

var (
	rxSubmatch   = regexp.MustCompile(Pair)
	rxReplace    = regexp.MustCompile(Trans)
	rxStarWithEe = regexp.MustCompile(StarWithEe)
)

func isBoolean(str string) bool {
	if _, err := strconv.ParseBool(str); err == nil {
		return true
	}
	return false
}

func replaceAllStringSubmatchFunc(str string) (string, string) {
	var key, value string
	for _, v := range rxSubmatch.FindAllSubmatchIndex([]byte(str), -1) {
		groups := []string{}
		for i := 0; i < len(v); i += 2 {
			groups = append(groups, str[v[i]:v[i+1]])
		}

		key = groups[1]
		value = groups[2]
	}

	return strings.Trim(key, " "), strings.Trim(value, " ")
}

func transform(line string) string {
	if line[0] == '[' {
		return "[" + rxReplace.ReplaceAllString(line[1:len(line)-1], "_") + "]"
	} else {
		key, value := replaceAllStringSubmatchFunc(line)

		if len(value) == 0 {
			value = fmt.Sprintf("\"\"")
		} else if valid.IsInt(value) {
			value = value
		} else if valid.IsFloat(value) {
			if rxStarWithEe.MatchString(value) {
				value = fmt.Sprintf("\"%v\"", value)
			} else {
				value = value
			}
		} else if isBoolean(value) {
			b, _ := strconv.ParseBool(value)
			value = strconv.FormatBool(b)
		} else {
			value = fmt.Sprintf("\"%s\"", value)
		}

		key = rxReplace.ReplaceAllString(key, "_")
		return fmt.Sprintf("%v=%v", key, value)
	}
}
