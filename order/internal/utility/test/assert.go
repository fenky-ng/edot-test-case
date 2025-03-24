package test

import "github.com/stretchr/testify/require"

func RequireErrorIs(expected error) require.ErrorAssertionFunc {
	return func(t require.TestingT, err error, msgAndArgs ...interface{}) {
		require.ErrorIs(t, err, expected, msgAndArgs...)
	}
}
