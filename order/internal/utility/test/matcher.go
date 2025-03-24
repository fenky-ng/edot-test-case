package test

import (
	"context"
	"reflect"

	"go.uber.org/mock/gomock"
)

var (
	ContextTypeMatcher = gomock.AssignableToTypeOf(reflect.TypeOf((*context.Context)(nil)).Elem())
)
