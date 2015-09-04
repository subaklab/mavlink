package mavlinkparse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXmlMarshal(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func OpenTestXml() Mavlink {
	xmlFile, err := os.Open("auto.xml")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer xmlFile.Close()

	XMLdata, _ := ioutil.ReadAll(xmlFile)

	var m Mavlink
	return m
	//xml.Unmarshal(XMLdata, &m)
}
