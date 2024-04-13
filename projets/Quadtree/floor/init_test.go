package floor

import (
	"testing"
)

func TestReadFloorFromFilePentacle(t *testing.T) {
	if ReadFloorFromFile("../floor-files/pentacle")[3][4] != 3 {
		t.Fail()
	}
}
func TestReadFloorFromFileBeaupasbeau(t *testing.T) {
	if ReadFloorFromFile("../floor-files/beaupasbeau")[3][4] != 0 {
		t.Fail()
	}
}
func TestReadFloorFromFileValidation(t *testing.T) {
	if ReadFloorFromFile("../floor-files/validation")[3][4] != 1 {
		t.Fail()
	}
}
func BenchmarkReadFloorFromFile(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		ReadFloorFromFile("../floor-files/pentacle")
	}
}
