package pointer

func PointerOf[T any](input T) (output *T) {
	output = &input
	return output
}

func ValueOf[T any](input *T) (output T) {
	if input != nil {
		output = *input
	}
	return output
}
