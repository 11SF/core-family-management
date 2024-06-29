package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/11SF/core-family-management/config"
	"github.com/11SF/core-family-management/pkg/v1/datamodel"
	"github.com/redis/go-redis/v9"
)

type familyRedis struct {
	rediskeys config.RedisKeys
	redis     *redis.Client
}

func NewFamilyRedis(redis *redis.Client, rediskeys config.RedisKeys) IFamilyRedis {
	return &familyRedis{redis: redis, rediskeys: rediskeys}
}
func (r *familyRedis) SaveFamily(ctx context.Context, family *datamodel.Family) error {

	key := fmt.Sprintf("%s:%s", r.rediskeys.Family, family.ID)

	byte, err := json.Marshal(family)
	if err != nil {
		return err
	}

	err = r.redis.Set(ctx, key, byte, 1*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *familyRedis) GetFamilyById(ctx context.Context, familyId string) (*datamodel.Family, error) {

	key := fmt.Sprintf("%s:%s", r.rediskeys.Family, familyId)

	resultByte, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	result := new(datamodel.Family)

	err = json.Unmarshal([]byte(resultByte), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *familyRedis) DeleteFamily(ctx context.Context, familyId string) error {

	key := fmt.Sprintf("%s:%s", r.rediskeys.Family, familyId)

	err := r.redis.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *familyRedis) SaveFamilyList(ctx context.Context, families []datamodel.Family, userId string) error {

	key := fmt.Sprintf("%s:%s", r.rediskeys.FamilyList, userId)

	byte, err := json.Marshal(families)
	if err != nil {
		return err
	}

	err = r.redis.Set(ctx, key, byte, 1*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *familyRedis) GetFamilyList(ctx context.Context, userId string) (*[]datamodel.Family, error) {

	key := fmt.Sprintf("%s:%s", r.rediskeys.FamilyList, userId)

	resultByte, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	result := new([]datamodel.Family)

	err = json.Unmarshal([]byte(resultByte), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *familyRedis) DeleteFamilyList(ctx context.Context, userId string) error {

	key := fmt.Sprintf("%s:%s", r.rediskeys.FamilyList, userId)

	err := r.redis.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
