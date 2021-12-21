package rng

func Deterministic() (func() uint64, func() uint64) {
	state := uint64(0)
	rolls := uint64(0)

	return func() uint64 {
			state++
			rolls++
			if state > 100 {
				state = 1
			}

			return state
		}, func() uint64 {
			return rolls
		}
}
