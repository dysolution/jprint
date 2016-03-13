package main

import (
	"strings"
	"testing"
)

func TestTryInt(t *testing.T) {
	var tests = []struct {
		given string
		want  int64
	}{
		{"2", int64(2)},
		{"0", int64(0)},
		{"-2147483648", int64(-2147483648)},
		{"2147483647", int64(2147483647)},
	}

	for _, tt := range tests {
		if got, err := tryInt(tt.given); got != tt.want {
			if err != nil {
				t.Errorf(`tryInt("%v") error: %v`, tt.given, err)
			}
			t.Errorf(`tryInt("%v") = %+v.(%v) want %v.(%T)`,
				tt.given, got, got, tt.want, tt.want)
		}
	}
}
func TestSimpleParse(t *testing.T) {
	var tests = []struct {
		given string
		want  interface{}
	}{
		{"2", int64(2)},
		{"1.0", float64(1.0)},
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		if got := simpleParse(tt.given); got != tt.want {
			t.Errorf(`parse("%v") = %+v.(%v) want %v.(%T)`,
				tt.given, got, got, tt.want, tt.want)
		}
	}
}

func TestSet(t *testing.T) {
	var tests = []struct {
		given string
		want  map[string]interface{}
	}{
		{"foo=bar", map[string]interface{}{"foo": "bar"}},
		{"bar=0", map[string]interface{}{"bar": int64(0)}},
	}
	p := params{}
	for _, tt := range tests {
		p.Set(tt.given)
		input := strings.Split(tt.given, "=")
		key, _ := input[0], input[1]

		got := p[key]
		want := tt.want[key]
		if got != tt.want[key] {
			t.Errorf(`p["%v"] = %v.(%T) want %v.(%T)`, key, got, got, want, want)
		}

		if got, ok := p["foo"]; !ok {
			t.Errorf(`parse("%v") = %+v.(%v) want %v.(%T)`,
				tt.given, got, got, tt.want, tt.want)
		}

	}
}
