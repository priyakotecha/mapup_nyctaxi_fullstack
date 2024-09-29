package queue

import (
	"context"
	"fmt"
	"nyctaxi_mapup/pkg/utils"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func AddToQueue(job string) error {
	err := redisClient.LPush(ctx, "jobs", job).Err()
	if err != nil {
		return fmt.Errorf("could not add job to queue: %v", err)
	}
	return nil
}

func ProcessQueue() error {
	for {
		job, err := redisClient.BRPop(ctx, 0, "jobs").Result()
		if err != nil {
			return fmt.Errorf("error processing queue: %v", err)
		}

		fmt.Printf("Processing job: %s\n", job[1])
		err = utils.ProcessParquetFile()
		if err != nil {
			return fmt.Errorf("error processing parquet file: %v", err)
		}
	}
}
