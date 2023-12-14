package rdb

import (
	"context"
	"encoding/json"
	"fmt"
	"frame/model"
	"time"
)

func (s *Store) formatUserKey(id uint) string {
	return fmt.Sprintf("user:%d", id)
}

func (s *Store) GetUsers(ctx context.Context, ids []uint) (map[uint]*model.User, error) {
	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = s.formatUserKey(id)
	}

	result, err := s.client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	recordMap := make(map[uint]*model.User)
	for i, data := range result {
		if data != nil {
			var record model.User
			if err := json.Unmarshal([]byte(data.(string)), &record); err != nil {
				return nil, err
			}

			recordMap[ids[i]] = &record
		}
	}

	return recordMap, nil
}

func (s *Store) CacheUsers(ctx context.Context, records []*model.User) error {
	pipe := s.client.Pipeline()

	for _, record := range records {
		key := s.formatUserKey(record.ID)
		data, err := json.Marshal(record)
		if err != nil {
			return err
		}
		pipe.Set(ctx, key, data, 5*time.Minute)
	}

	_, err := pipe.Exec(ctx)
	return err
}

func (s *Store) ClearUser(ctx context.Context, id uint) error {
	return s.client.Del(ctx, s.formatUserKey(id)).Err()
}
