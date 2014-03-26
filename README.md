gopencorpora
============

Parsed .xml dictionary from opencorpora.org

example xml file:

```xml
<dictionary version="0.8" revision="403605">
    <grammemes>
        <grammeme parent="">POST</grammeme>
        <grammeme parent="POST">NOUN</grammeme>
    </grammemes>
    <lemmata>
        <lemma id="1" rev="402007">
            <l t="абажур"><g v="NOUN"/><g v="inan"/><g v="masc"/></l>
            <f t="абажур"><g v="sing"/><g v="nomn"/></f>
            <f t="абажура"><g v="sing"/><g v="gent"/></f>
        </lemma>
    </lemmata>
    <link_types>
        <type id="1">VERB_GERUND</type>
    </link_types>
    <links>
        <link id="1" from="104" to="106" type="1"/>
    </links>
</dictionary>
```

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