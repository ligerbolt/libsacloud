package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testCDROMName = "libsacloud_test_iso_image"

func TestCRUDCDROM(t *testing.T) {
	api := client.CDROM

	//CREATE
	newCD := api.New()
	newCD.Name = testCDROMName
	newCD.Description = "hoge"
	newCD.SizeMB = 5120

	cd, _, err := api.Create(newCD)

	assert.NoError(t, err)
	assert.NotEmpty(t, cd)
	id := cd.ID

	//READ
	cd, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, cd)

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)
}

func init() {
	testSetupHandlers = append(testSetupHandlers, cleanupCDROM)
	testTearDownHandlers = append(testTearDownHandlers, cleanupCDROM)
}

func cleanupCDROM() {
	items, _ := client.CDROM.Reset().WithNameLike(testCDROMName).Find()
	for _, item := range items.CDROMs {
		client.CDROM.Delete(item.ID)
	}
}
