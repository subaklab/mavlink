package goparse

import (
	"encoding/xml"
)

type Mavlink struct {
	XMLName  xml.Name `xml:"mavlink"`
	include  string   `xml:"include"`
	Version  string   `xml:"version"`
	Enums    []enums
	Messages []messages
}

type enums struct {
	XMLName  xml.Name `xml:"enums"`
	EnumList []enum   `xml:"enum"`
}

type enum struct {
	name        string `xml:"name,attr"`
	Description string `xml:"description"`
	Entry       entry
}

type entry struct {
	XMLName     xml.Name `xml:"entry,attr"`
	value       string   `xml:"value,attr"`
	value       name     `xml:"name,attr"`
	Description string   `xml:"description"`
}

type messages struct {
	XMLName     xml.Name  `xml:"messages"`
	MessageList []message `xml:"message"`
}

type message struct {
	id          string `xml:"id,attr"`
	name        string `xml:"name,attr"`
	description string
	Field       field `xml:"field"`
}

type field struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
	name string `xml:"name,attr"`
}
