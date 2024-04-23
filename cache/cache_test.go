package cache_test

import (
	"testing"

	"github.com/george-e-shaw-iv/go/cache"
	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	const capacity = 2

	tt := []struct {
		Name           string
		Implementation cache.Cache[string, string]
	}{
		{
			Name:           "LRU",
			Implementation: cache.NewLRU[string, string](capacity),
		},
	}

	for _, test := range tt {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			_, err := test.Implementation.Get("foobar")
			assert.ErrorIs(t, err, cache.ErrNotInCache)

			test.Implementation.Put("foo", "bar")
			foo, err := test.Implementation.Get("foo")
			assert.NoError(t, err)
			assert.Equal(t, "bar", foo)

			test.Implementation.Put("bar", "baz")
			bar, err := test.Implementation.Get("bar")
			assert.NoError(t, err)
			assert.Equal(t, "baz", bar)

			test.Implementation.Put("george", "shaw")
			george, err := test.Implementation.Get("george")
			assert.NoError(t, err)
			assert.Equal(t, "shaw", george)

			// Capacity was exceeding on last put, so the last entry should've been expelled.
			_, err = test.Implementation.Get("foo")
			assert.ErrorIs(t, err, cache.ErrNotInCache)
		})
	}
}
