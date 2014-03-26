// Parse .xml dictionary from opencorpora.org
package gopencorpora

import (
	"io"
	"encoding/xml"
)

type Dictionary struct {
	Version   float32 `xml:"version,attr"`
	Revision  int `xml:"revision,attr"`
	Grammemes []*Grammeme `xml:"grammemes>grammeme"`
	Lemmata   []*Lemma `xml:"lemmata>lemma"`
	LinkTypes []*Type `xml:"link_types>type"`
	Links     []*Link `xml:"links>link"`
}

type Grammeme struct {
	Name   string `xml:",chardata"`
	Parent string `xml:"parent,attr"`
}

type Lemma struct {
	Id       int `xml:"id,attr"`
	Revision int `xml:"rev,attr"`
	Default *LemmaDefault `xml:"l"`
	Forms    []*LemmaForm `xml:"f"`
}

type LemmaDefault struct {
	Name     string `xml:"t,attr"`
	Grammems []LemmaGrammem `xml:"g"`
}
type LemmaForm struct {
	Name     string `xml:"t,attr"`
	Grammems []LemmaGrammem `xml:"g"`
}
type LemmaGrammem struct {
	Grammem string `xml:"v,attr"`
}

type Type struct {
	Id   int `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Link struct {
	Id   int `xml:"id,attr"`
	From int `xml:"from,attr"`
	To   int `xml:"to,attr"`
	Type int `xml:"type,attr"`
}

func (dict *Dictionary) Parse(reader io.Reader) error {
	decoder := xml.NewDecoder(reader)
	err := decoder.Decode(dict)
	return err
}
