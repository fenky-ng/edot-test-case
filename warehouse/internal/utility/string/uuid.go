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
