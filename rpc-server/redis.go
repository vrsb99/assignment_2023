package main

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type DBClient struct {
	redisClient *redis.Client
}

func (dbClient *DBClient) StartDatabase(ctx context.Context) error {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "TikTok",
		DB:       0,
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		return err
	}

	dbClient.redisClient = redisClient
	return nil
}

func (dbClient *DBClient) WriteToDatabase(ctx context.Context, id string, input *Input) error {
	jsonMessage, err := json.Marshal(input)
	if err != nil {
		return err
	}

	zsetMember := &redis.Z{
		Score:  float64(input.Timestamp),
		Member: jsonMessage,
	}

	if _, err = dbClient.redisClient.ZAdd(ctx, id, *zsetMember).Result(); err != nil {
		return err
	}

	return nil
}

func (dbClient *DBClient) ReadFromDatabase(ctx context.Context, id string, start int64, end int64, reverse bool) ([]*Input, error) {
	var rawMessages []string
	var err error

	if reverse {
		rawMessages, err = dbClient.redisClient.ZRevRange(ctx, id, start, end).Result()
	} else {
		rawMessages, err = dbClient.redisClient.ZRange(ctx, id, start, end).Result()
	}

	if err != nil {
		return nil, err
	}

	processedMessages := make([]*Input, 0, len(rawMessages))
	for _, rawMessage := range rawMessages {
		message := &Input{}
		if err := json.Unmarshal([]byte(rawMessage), message); err != nil {
			return nil, err
		}

		processedMessages = append(processedMessages, message)
	}

	return processedMessages, nil
}
