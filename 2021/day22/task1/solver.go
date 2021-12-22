package task1

import (
	"aoc-2021-day22/geometry"
	"aoc-2021-day22/input"
)

func Solve(data input.Data) (uint64, error) {
	filtered := filter(data)
	cuboids := make([]geometry.Cuboid, 0)

	for _, op := range filtered.Operations {
		cuboidsAfterOp := make([]geometry.Cuboid, 0)

		for _, existing := range cuboids {
			if notIntersect(op.Cuboid, existing) {
				cuboidsAfterOp = append(cuboidsAfterOp, existing)
				continue
			}

			if op.Cuboid.LowerX > existing.LowerX {
				cube := geometry.Cuboid{
					LowerX: existing.LowerX,
					UpperX: op.Cuboid.LowerX - 1,
					LowerY: existing.LowerY,
					UpperY: existing.UpperY,
					LowerZ: existing.LowerZ,
					UpperZ: existing.UpperZ,
				}

				cuboidsAfterOp = append(cuboidsAfterOp, cube)
			}

			if op.Cuboid.UpperX < existing.UpperX {
				cube := geometry.Cuboid{
					LowerX: op.Cuboid.UpperX + 1,
					UpperX: existing.UpperX,
					LowerY: existing.LowerY,
					UpperY: existing.UpperY,
					LowerZ: existing.LowerZ,
					UpperZ: existing.UpperZ,
				}

				cuboidsAfterOp = append(cuboidsAfterOp, cube)
			}

			if op.Cuboid.LowerY > existing.LowerY {
				cube := geometry.Cuboid{
					LowerX: max(op.Cuboid.LowerX, existing.LowerX),
					UpperX: min(op.Cuboid.UpperX, existing.UpperX),
					LowerY: existing.LowerY,
					UpperY: op.Cuboid.LowerY - 1,
					LowerZ: existing.LowerZ,
					UpperZ: existing.UpperZ,
				}

				cuboidsAfterOp = append(cuboidsAfterOp, cube)
			}

			if op.Cuboid.UpperY < existing.UpperY {
				cube := geometry.Cuboid{
					LowerX: max(op.Cuboid.LowerX, existing.LowerX),
					UpperX: min(op.Cuboid.UpperX, existing.UpperX),
					LowerY: op.Cuboid.UpperY + 1,
					UpperY: existing.UpperY,
					LowerZ: existing.LowerZ,
					UpperZ: existing.UpperZ,
				}

				cuboidsAfterOp = append(cuboidsAfterOp, cube)
			}

			if op.Cuboid.LowerZ > existing.LowerZ {
				cube := geometry.Cuboid{
					LowerX: max(op.Cuboid.LowerX, existing.LowerX),
					UpperX: min(op.Cuboid.UpperX, existing.UpperX),
					LowerY: max(op.Cuboid.LowerY, existing.LowerY),
					UpperY: min(op.Cuboid.UpperY, existing.UpperY),
					LowerZ: existing.LowerZ,
					UpperZ: op.Cuboid.LowerZ - 1,
				}

				cuboidsAfterOp = append(cuboidsAfterOp, cube)
			}

			if op.Cuboid.UpperZ < existing.UpperZ {
				cube := geometry.Cuboid{
					LowerX: max(op.Cuboid.LowerX, existing.LowerX),
					UpperX: min(op.Cuboid.UpperX, existing.UpperX),
					LowerY: max(op.Cuboid.LowerY, existing.LowerY),
					UpperY: min(op.Cuboid.UpperY, existing.UpperY),
					LowerZ: op.Cuboid.UpperZ + 1,
					UpperZ: existing.UpperZ,
				}

				cuboidsAfterOp = append(cuboidsAfterOp, cube)
			}
		}

		if op.On {
			cuboidsAfterOp = append(cuboidsAfterOp, op.Cuboid)
		}

		cuboids = cuboidsAfterOp
	}

	totalVol := uint64(0)
	for _, x := range cuboids {
		totalVol += uint64(x.Volume())
	}

	return totalVol, nil
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

func notIntersect(c1, c2 geometry.Cuboid) bool {
	return c1.LowerX > c2.UpperX ||
		c1.UpperX < c2.LowerX ||
		c1.LowerY > c2.UpperY ||
		c1.UpperY < c2.LowerY ||
		c1.LowerZ > c2.UpperZ ||
		c1.UpperZ < c2.LowerZ
}

func filter(data input.Data) input.Data {
	operations := make([]geometry.Operation, 0)

	for _, v := range data.Operations {
		if valid(v.Cuboid.LowerX) &&
			valid(v.Cuboid.UpperX) &&
			valid(v.Cuboid.LowerY) &&
			valid(v.Cuboid.UpperY) &&
			valid(v.Cuboid.LowerZ) &&
			valid(v.Cuboid.UpperZ) {
			operations = append(operations, v)
		}
	}

	return input.Data{Operations: operations}
}

func valid(num int64) bool {
	return num >= -50 && num <= 50
}
