# jprint
jprint is a simple CLI JSON generator inspired by [jo](https://github.com/jpmens/jo). It takes an set of key/value pairs in the format `key=value` and prints out a corresponding object in JSON.

# Examples

Per the JSON spec, keys are strings. All values are strings by default.

```bash
$ jprint foo=bar
{
  "foo": "bar"
}
```

Integers and floating-point numbers are detected and converted into 64-bit representations, and boolean values are autodetected:

```bash
$ jprint tau=6.283185 right_out=5 proprietary=false
{
  "proprietary": false,
  "right_out": 5,
  "tau": 6.283185
}
```

jprint uses Go's `json.MarshalIndent`, which sorts keys alphabetically by default:

```bash
$ jprint foo=bar num=3 a_few_numbers=[1,2,3]
{
  "a_few_numbers": [
    1,
    2,
    3
  ],
  "foo": "bar",
  "num": 3
}
```

# Installation

```bash
$ git clone https://github.com/dysolution/jprint.git
$ cd jprint
$ go install
```

# Tests

```bash
$ ./test.sh
```

# TODO

- [ ] support nesting, e.g., `jprint bar=$(jprint foo=3)`

# License

MIT
