// +build !race

package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordSuccess(t *testing.T) {
	password := CreatePassword("secret")
	assert.Equal(t, true, ComparePassword(password, []byte("secret")))
}

func TestPasswordFailure(t *testing.T) {
	password := CreatePassword("secret")
	assert.Equal(t, false, ComparePassword(password, []byte("secretx")))
}

func TestBCryptFailure(t *testing.T) {
	assert.Panics(t, func() { CreatePassword("secret") })
}
