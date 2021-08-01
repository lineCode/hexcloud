package hexgrid

type Hexagon struct {
	X int64
	Y int64
	Z int64
}

type HexCoord struct {
	Q int64
	R int64
}

var Directions = []Hexagon {
	NewHexQR(1,0 ),
	NewHexQR(1,-1),
	NewHexQR(0,-1),
	NewHexQR(-1,0),
	NewHexQR(-1,1),
	NewHexQR(0,1),
}

func NewHexQR(q, r int64) Hexagon {
	return Hexagon {
		X:         q,
		Y:         r,
		Z:         -q - r,
	}
}

func NewHexXYZ(x, y, z int64) Hexagon {
	return Hexagon {
		X:         x,
		Y:         y,
		Z:         z,
	}
}

func Add(h1 , h2 Hexagon) Hexagon {
	return Hexagon{
		X: h1.X + h2.X,
		Y: h1.Y + h2.Y,
		Z: h1.Z + h2.Z,
	}
}

func Scale(hc Hexagon, radius int64) Hexagon {
	return Hexagon{
		X:         hc.X * radius,
		Y:         hc.Y * radius,
		Z:         hc.Z * radius,
	}
}

func Neighbour(hexagon Hexagon, direction int) Hexagon {
	return Add(hexagon, Directions[direction])
}

func Ring(center Hexagon, radius int64) []Hexagon {
	hexes := []Hexagon{}
	hexagon := Add(center,Scale(Directions[4], radius))

	for i := 0; i < 6; i++ {
		for j := 0; j < int(radius); j++ {
			hexes = append(hexes, hexagon)
			hexagon = Neighbour(hexagon, i)
		}
	}

	return hexes
}
