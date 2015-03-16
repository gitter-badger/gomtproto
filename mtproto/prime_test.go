package mtproto_test

import (
	"github.com/onrik/gomtproto/mtproto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSimple(t *testing.T) {
	assert.Equal(t, true, mtproto.IsPrime(1))
	assert.Equal(t, false, mtproto.IsPrime(2))
	assert.Equal(t, false, mtproto.IsPrime(144))

	assert.Equal(t, false, mtproto.IsPrime(144))
	assert.Equal(t, true, mtproto.IsPrime(4091))
	assert.Equal(t, true, mtproto.IsPrime(1738498649))
	assert.Equal(t, true, mtproto.IsPrime(1073753687))
	assert.Equal(t, true, mtproto.IsPrime(1302848473))
}

func TestFactorize(t *testing.T) {
	assert.Equal(t, 3, mtproto.Factorize(21))
	assert.Equal(t, 23, mtproto.Factorize(2323))

	assert.Equal(t, 1481380559, mtproto.Factorize(2265400785806246213))
	assert.Equal(t, 1738498649, mtproto.Factorize(3054904118220193429))
}
