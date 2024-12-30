package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// Define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// Top level declarations for the storeService and Redis context
var (
	storeService = &StorageService{} // Create a pointer to a StorageService instance
    ctx = context.Background() // Create a root context
)

// inshAllah in the future this should be changed to an LRU
const CacheDuration = 6 * time.Hour

// Initializing the store service and return a store pointer 
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options { // creates a new redisClient
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}


/* We want to be able to save the mapping between the originalUrl 
and the generated shortUrl url
*/

func SaveUrlMapping(shortUrl string, originalUrl string, userId string){ 
	// .Set() -> Sets the string value of a key, ignoring its type. The key is created if it doesn't exist.
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

/*
We should be able to retrieve the initial long URL once the short 
is provided. This is when users will be calling the shortlink in the 
url, so what we need to do here is to retrieve the long url and
think about redirect.
*/

func RetrieveInitialUrl(shortUrl string) string {
	// Use the Redis client to attempt to retrieve the value (long URL) associated with the given short URL key.
    // The `Get` method sends a GET request to the Redis server, and returns a redis command object
	// The Result method executes the redis command and retrieves the result, and returns the result and error if one exists
    // `ctx` provides context, which can include things like timeouts or cancellation signals for the request.
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()

	// If there's an error (e.g., the key does not exist in Redis), panic and print an error message.
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
