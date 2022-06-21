# Golang References

## Effective Go

https://go.dev/doc/effective_go

## Format

https://pkg.go.dev/fmt

### Formatting errors

Always use `%w` to format errors, instead of using `%s` or `%v`.  The idea is error messages can be chained, thus
wrapped or unwrapped.  Sample usage: https://gosamples.dev/wrap-unwrap-errors/





