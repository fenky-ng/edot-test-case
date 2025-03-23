package string

import "github.com/google/uuid"

func ParseStringArrToUuidArr(in []string) (out []uuid.UUID, err error) {
	for _, data := range in {
		cur, err := uuid.Parse(data)
		if err != nil {
			return out, err
		}
		out = append(out, cur)
	}
	return out, nil
}

func ParseUuidArrToStringArr(in []uuid.UUID) (out []string) {
	for _, data := range in {
		out = append(out, data.String())
	}
	return out
}
