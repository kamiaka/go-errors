# go-errors

## Adding labeled value to an error

```go
_, err := Foo(a, b) // returned fmt.Errorf("Some error")
if err != nil {
	return errors.WithFields(err).String("a", a).Int("b", b)
}
```

```
Some error
	a:value of a\\n	b:42
```

## LICENSE

MIT