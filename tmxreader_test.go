package tmxreader

import (
	"testing"
)

var data = `
<?xml version="1.0" encoding="UTF-8"?>
<map version="1.0" orientation="orthogonal" width="50" height="50" tilewidth="8" tileheight="8">
 <properties>
  <property name="id" value="test"/>
  <property name="group" value="testgroup"/>
 </properties>
 <tileset firstgid="1" name="world.tiles" tilewidth="8" tileheight="8">
  <image source="../images/world.tiles.png" width="768" height="512"/>
 </tileset>
 <tileset firstgid="6145" name="character.tiles" tilewidth="8" tileheight="8">
  <image source="../images/character.tiles.png" width="400" height="300"/>
 </tileset>
 <layer name="Tile Layer 1" width="50" height="50">
  <data encoding="csv">
  	layer1,data1
  </data>
 </layer>
 <layer name="Tile Layer 2" width="50" height="50">
  <data encoding="csv">
  	layer1,data1
  </data>
 </layer>
 <objectgroup name="Object Layer 1" width="50" height="50">
  <object type="Object 1" x="0" y="19" width="375" height="51"/>
  <object type="Object 2" x="16" y="88" width="46" height="21"/>
 </objectgroup>
</map>`

var malformedXMLData = `
<?xml version="1.0" encoding="UTF-8"?>
<map version="1.0">
</mop>`

func TestParse(t *testing.T) {
	tmxmap, err := Parse([]byte(data))
	if err != nil {
		t.Fatal("Failed to parse XML")
	} else {
		if tmxmap.Version != "1.0" {
			t.Log("Incorrect version number")
		}
		if len(tmxmap.Properties) != 1 {
			logNodeCountError(t, "Properties")
		} else if len(tmxmap.Properties[0].Property) != 2 {
			logNodeCountError(t, "Property")
		}
		if len(tmxmap.Tilesets) != 2 {
			logNodeCountError(t, "Tilesets")
		} else if len(tmxmap.Tilesets[0].Images) != 1 {
			logNodeCountError(t, "Images")
		}
		if len(tmxmap.Layers) != 2 {
			logNodeCountError(t, "Layers")
		}
		if len(tmxmap.ObjectGroups) != 1 {
			logNodeCountError(t, "ObjectGroups")
		} else if len(tmxmap.ObjectGroups[0].Objects) != 2 {
			logNodeCountError(t, "Object")
		}
	}
}

func TestParseMalformedXML(t *testing.T) {
	_, err := Parse([]byte(malformedXMLData))
	if err == nil {
		t.Fatal("Malformed XML parsed successfully")
	}
}

func logNodeCountError(t *testing.T, nodeName string) {
	t.Logf("Failed to read node: %s", nodeName)
	t.Fail()
}
