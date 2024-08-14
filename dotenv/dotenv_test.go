package dotenv

import (
	// "os"
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_Load(t *testing.T) {
	
	t.Run("testing no error env file", func(t *testing.T) {
		err := Load()
		assert.Equal(t, nil, err)
		
	})

	t.Run("testing non existing file", func(t *testing.T) {
		err := Load("nonexisting.env")
		assert.NotEqual(t, nil, err)
	})

	t.Run("testing incorrect env ile", func(t *testing.T) {
		err := Load("bad.env")
		assert.NotEqual(t, nil, err)
	})
}