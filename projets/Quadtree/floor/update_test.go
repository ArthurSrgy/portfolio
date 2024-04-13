package floor

import "testing"

func TestUpdateFromFileFloor(t *testing.T) {
	var f = Floor{content: nil,
		fullContent: ReadFloorFromFile("../floor-files/pentacle")}
	f.UpdateFromFileFloor(0, 0)
	if f.fullContent[0][0] != f.content[len(f.content)/2][len(f.content[0])/2] || f.content[len(f.content)/2-1][len(f.content[0])/2-1] != -1 {
		t.Fail()
	}
}
