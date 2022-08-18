package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilder(t *testing.T) {
	builder := NewBuilder()
	assert.Equal(t, 1, builder.Build().id)

	builder2 := NewBuilder()
	id := builder2.Clone().id
	assert.Equal(t, 1, id)

}
