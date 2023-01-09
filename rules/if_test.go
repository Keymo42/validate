package rules

import (
	"testing"

	"github.com/jdtron/validate"
	"github.com/stretchr/testify/assert"
)

func TestIfNotEmpty(t *testing.T) {
	tt := []struct {
		name      string
		condition bool
		value     string
		expectOK  bool
	}{
		{name: "run valid", value: "not-empty", expectOK: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ok := validate.New().
				With(
					tc.name,
					If(tc.condition, func() validate.FieldValidator[string] {
						return Str(tc.value).NotEmpty()
					}),
				).
				Run().
				OK()
			assert.Equal(t, tc.expectOK, ok)
		})
	}
}
