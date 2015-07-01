package cache

import (
	"cache"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SetData(t *testing.T) {
	c := cache.New("goleveldb", false)
	for i := 0; i < 1000; i++ {
		key_str := fmt.Sprint("key", i)
		val_str := fmt.Sprint("value", i)
		err := c.SetData(key_str, val_str, 1000000)
		assert.Nil(t, err, "Hope we don't have a error")
	}
	c.CloseConnection()
}

func Test_GetData(t *testing.T) {
	c := cache.New("goleveldb", false)
	x, err := c.GetData("key999")
	assert.Nil(t, err, "no error")
	assert.Equal(t, "value999", x)
	c.CloseConnection()
}

func Test_DeleteData(t *testing.T) {
	c := cache.New("goleveldb", false)
	err := c.DeleteData("key999")
	assert.Nil(t, err, "no error")
	c.CloseConnection()
}

func Test_UpdateTTL(t *testing.T) {
	c := cache.New("goleveldb", false)
	_, err := c.UpdateTTL("key999", 20000)
	assert.Nil(t, err, "no error")
	c.CloseConnection()

}
