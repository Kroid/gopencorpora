gopencorpora
============

Parsed .xml dictionary from opencorpora.org

usage:

```go
var fileName string = "opencorpora-dict.xml"
file, err := os.Open(fileName)
if err != nil {
  fmt.Printf("error:n%s\n")
  os.Exit(1)
}

dict := &opencorpora.Dictionary{}
if err = dict.Parse(file); err != nil {
  fmt.Printf("xml parse error:\n%s\n", err)
}

fmt.Printf("%s\n", dict)
```