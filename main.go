package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	flag.Var(paramsFlags, "p", "key/value pair, e.g.: foo=bar")
}

type params map[string]interface{}

var paramsFlags = make(params)

func (p params) String() string {
	out, _ := json.MarshalIndent(p, "", "  ")
	// out, _ := json.Marshal(p)
	return fmt.Sprintf("%s", out)
}

func (p params) Set(rawKV string) error {
	k, v, err := parseKV(rawKV)
	if err != nil {
		return err
	}
	p[k] = simpleParse(v)
	return nil
}

func parseKV(input string) (string, string, error) {
	tokens := strings.Split(input, "=")
	key := tokens[0]
	if len(tokens) < 2 {
		return key, "", errors.New("missing value for '" + tokens[0] + "'")
	}
	val := tokens[1]
	return key, val, nil
}

func (p params) parseArgs(input []string) {
	for _, param := range input {
		val := simpleParse(param).(string)
		p.Set(val)
	}
}

func tryFloat(val string) (float64, error) {
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0.0, errors.New("can't convert to float64")
	}
	return f, nil
}

func tryInt(val string) (int64, error) {
	f, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, errors.New("can't convert to int64")
	}
	return f, nil
}

func simpleParse(val string) interface{} {
	i, err := strconv.ParseInt(val, 10, 64)
	if err == nil {
		return i
	}
	f, err := strconv.ParseFloat(val, 64)
	if err == nil {
		return f
	}
	b, err := strconv.ParseBool(val)
	if err == nil {
		return b
	}

	var j []interface{}
	err = json.Unmarshal([]byte(val), &j)
	if err == nil {
		return j
	}

	var aj interface{}
	bytes := []byte(val)
	err = json.Unmarshal(bytes, &aj)
	if &aj == nil {
		m := aj.(map[string]interface{})
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")

				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}
		if err == nil {
			return aj
		}
	}

	return val
}

func main() {
	paramsFlags.parseArgs(os.Args[1:])
	flag.Parse()
	fmt.Print(paramsFlags)
}
