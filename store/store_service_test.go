package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertAndRetrieval(t *testing.T) {
	initialUrl := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUID := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortUrl := "Jsz4k57oAX"

	// Persist data mapping
	SaveURLMapping(shortUrl, initialUrl, userUUID)

	retrievedUrl := RetrieveInitialURL(shortUrl)

	assert.Equal(t, initialUrl, retrievedUrl)
}