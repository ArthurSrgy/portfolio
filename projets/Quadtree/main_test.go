package floor

import (
	"testing"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

func BenchmarkReadFloorFromFile(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		ReadFloorFromFile(configuration.Global.FloorFile)
	}
}
