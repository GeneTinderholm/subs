package shift

import (
	"github.com/GeneTinderholm/cmf"
	"github.com/GeneTinderholm/cmf/this"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDo(t *testing.T) {
	// TODO - this is a placeholder so go mod tidy doesn't clean things up
	bs := cmf.Must(os.ReadFile(this.Dir() + "../sub.srt"))
	assert.Len(t, bs, 100)
}
