package redis

import (
	"context"
	"encoding/json"
	"fmt"

	redisclient "github.com/crypto-tester/terraform-backend/pkg/client/redis"

	"github.com/crypto-tester/terraform-backend/pkg/terraform"
)

var ctx = context.Background()

const Name = "redis"

type RedisStorage struct {
	db    int
	table string
}

func NewRedisStorage(table string) (*RedisStorage, error) {
	rdb, err := redisclient.NewClient()

	if err != nil {
		return nil, fmt.Errorf("не удалось сериализовать состояние: %w", err)
	}

	p := &RedisStorage{
		db:    0,
		table: table,
	}

	err = rdb.Set(ctx, table, 0, 0).Err()

	return p, nil
}

func SaveState(s *terraform.State) error {
	rdb, err := redisclient.NewClient()

	if err != nil {
		return fmt.Errorf("не удалось сериализовать состояние: %w", err)
	}

	stateJSON, err := json.Marshal(s.Data)
	if err != nil {
		return fmt.Errorf("не удалось сериализовать состояние: %w", err)
	}

	err = rdb.Set(ctx, s.ID, stateJSON, 0).Err()

	if err != nil {
		return fmt.Errorf("не удалось сохранить состояние в Redis: %w", err)
	}
	return nil
}

func GetName() string {
	return Name
}

// func (p *RedisStorage) SaveState(s *terraform.State) error {
// 	if _, err := p.rdb.Set(ctx, s.ID, s.Data); err != nil {

// 		return err
// 	}

// 	return nil
// }

// func (p *RedisStorage) GetState(id string) (*terraform.State, error) {
// 	s := &terraform.State{}

// 	err := p.rdb.Get(ctx, id).Result()
// 	if err != nil {
// 		return nil, storage.ErrStateNotFound
// 	} else if err != nil {
// 		return nil, err
// 	}

// 	return s, nil
// }

// func (p *RedisStorage) DeleteState(id string) error {
// 	return p.rdb.Del(ctx, id).Result()

// 	if err != nil {
// 		panic(err)
// 	}
// }
