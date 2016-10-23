# timerange

timerange is a library to parse string with time ranges in them into
`time.Time`.

## Examples

It works with time ranges such as:

```go
values, err := timerange.Parse("2016/10/22..2016/10/24")
// []time.Time{
//    2016-10-22 00:00:00 +0000 UTC
//    2016-10-23 00:00:00 +0000 UTC
//    2016-10-24 00:00:00 +0000 UTC
// }
```

And also individual timestamps:

```go
value, err := timerange.Parse("2016/10/22")
// []time.Time{2016-10-22 00:00:00 +0000 UTC}
```

In addition it implements `flag.Value` so it can be used by the `flag` package.

```go
t := timerange.NewTimerange()
flag.Var(&t, "timestamp", "Pass single or range of timestamps")
```

The default layout for parsing time and range seperator are: `2006/01/02` and
`..`, however that can be customized by initializing your own `Timerange`.

```go
t := timerange.Timerange{
TimeValues :    []time.Time{},
           TimeLayout :    "2016-01-02",
           RangeSeparator: "..",
}
flag.Var(&t, "timestamp", "Pass single or range of timestamps")
```

## Test

```
go get github.com/smartystreets/goconvey/convey
go test
```
