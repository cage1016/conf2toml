package convert

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	Int   string = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	Float string = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
	Pair  string = "(.*) =(.*)"
	Trans string = "[[:punct:]]|[[:space:]]"
)

var (
	rxInt      = regexp.MustCompile(Int)
	rxFloat    = regexp.MustCompile(Float)
	rxSubmatch = regexp.MustCompile(Pair)
	rxReplace  = regexp.MustCompile(Trans)
)

func isInt(str string) bool {
	if len(str) == 0 {
		return true
	}
	return rxInt.MatchString(str)
}

func isFloat(str string) bool {
	return str != "" && rxFloat.MatchString(str)
}

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

	return key, strings.Trim(value, " ")
}

func Transform(line string) string {
	if line[0] == '[' {
		return "[" + rxReplace.ReplaceAllString(line[1:len(line)-1], "_") + "]"
	} else {
		key, value := replaceAllStringSubmatchFunc(line)
		if len(value) == 0 {
			value = fmt.Sprintf("\"\"")
		} else if isInt(value) || isFloat(value) {
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
