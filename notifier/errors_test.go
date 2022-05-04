package notifier

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageError(t *testing.T) {
	const expectedErrString = "test err"
	baseErr := errors.New(expectedErrString)
	err := &MessageError{
		Err: baseErr,
	}

	assert.Equal(t, expectedErrString, err.Error())
	assert.True(t, errors.Is(err, baseErr))
}
