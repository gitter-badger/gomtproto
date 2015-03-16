package mtproto_test

import (
	"github.com/onrik/gomtproto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchema(t *testing.T) {
	err := mtproto.LoadSchema("schema.json")
	assert.Equal(t, nil, err)
	assert.Equal(t, 122, len(mtproto.TLSchema.Methods))
	assert.Equal(t, 339, len(mtproto.TLSchema.ConstructorsId))
	assert.Equal(t, 140, len(mtproto.TLSchema.ConstructorsType))
}
