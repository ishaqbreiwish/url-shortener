package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// initializing empty storage service for testing
var testStoreService = &StorageService{}

// intializes redis client using initializeStore()
func init() {
	testStoreService = InitializeStore()
}

// Use Test in the function name so that Go knows its a test case
// t is a pointer to the testing.T type
// The testing.T object is used to log errors or failures in test cases and to control test execution
func TestStoreInit(t *testing.T) {
	// Use the `assert` library to check a condition in a test case.
    // Here, we're asserting that `testStoreService.redisClient` is not nil.
	assert.True(t, testStoreService.redisClient != nil)
}

// Tests Insertion and retrieval
func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	// Persist data mapping
	SaveUrlMapping(shortURL, initialLink, userUUId)

	// Retrieve initial URL
	retrievedUrl := RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialLink, retrievedUrl)
}
