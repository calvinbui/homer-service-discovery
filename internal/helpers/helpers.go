package helpers

import "strconv"

func String(v string) *string {
	return &v
}

func StringToFileMode(v string) (uint32, error) {
	s := "0" + v

	f64, err := strconv.ParseUint(s, 2, 32)
	if err != nil {
		return 0, err
	}

	f32 := uint32(f64)

	return f32, nil
}
