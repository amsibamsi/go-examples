package intersect

const (
	nearZero = 1e-10
)

type vec struct {
	x, y, z float64
}

func dot(v, w *vec) float64 {
	return v.x*w.x + v.y*w.y + v.z*w.z
}

func (v *vec) sub(w *vec) *vec {
	return &vec{
		x: v.x - w.x,
		y: v.y - w.y,
		z: v.z - w.z,
	}
}

type line struct {
	start, end vec3
}

type tri struct {
	a, b, c vec3
	n       vec3
}

// - check if perpendicular: scalar of t normal with vector on l is near 0
//   - yes: check if distance is near 0
//     - yes: intersect l with t in 2d (how?)
//     - no: perpendicular, return nil
//   - no: intersect l with plane that includes t
//     - check if intersection lies within t (comcpute right/left of for every
//       triangle line?)
func intersect(t *tri, l *line) *vec {
	p := sub(&l.end, &l.start)
	if dot(p, &t.n) <= nearZero {
	}
}
