package test

import (
	"context"
	"fmt"
	"go_category/configs"
	"go_category/domain/db"
	"testing"
)

func TestRedisIncr(t *testing.T) {
	ctx := context.Background()
	test(ctx, t, configs.REDIS_TYPE, func() {
		rds := db.RedisClient
		r, err := rds.Incr(ctx, "wenzhouyao_use_cnt").Result()
		if err != nil {
			fmt.Println("rds.Incr fail, err:", err.Error())
			return
		}
		fmt.Println("rds.Incr reply:", r)
		cnt, err := rds.Get(ctx, "wenzhouyao_use_cnt").Int()
		if err != nil {
			fmt.Println("rds.Get fail, err:", err.Error())
		} else {
			fmt.Println("wenzhouyao_use_cnt:", cnt)
		}
	})
}

func TestRedisIncrBy(t *testing.T) {
	ctx := context.Background()
	test(ctx, t, configs.REDIS_TYPE, func() {
		rds := db.RedisClient
		r, err := rds.IncrBy(ctx, "wenzhouyao_use_cnt", 10).Result()
		if err != nil {
			fmt.Println("rds.IncrBy fail, err:", err.Error())
			return
		}
		fmt.Println("rds.IncrBy reply:", r)
		cnt, err := rds.Get(ctx, "wenzhouyao_use_cnt").Int()
		if err != nil {
			fmt.Println("rds.Get fail, err:", err.Error())
		} else {
			fmt.Println("wenzhouyao_use_cnt:", cnt)
		}
	})
}

func TestRedisSet(t *testing.T) {
	ctx := context.Background()
	test(ctx, t, configs.REDIS_TYPE, func() {
		rds := db.RedisClient
		// 参数expiration为0表示此值没有过期时间
		r, err := rds.Set(ctx, "wenzhouyao_use_cnt", 100, 0).Result()
		if err != nil {
			fmt.Println("rds.Set fail, err:", err.Error())
			return
		}
		fmt.Println("rds.Set reply:", r)
		cnt, err := rds.Get(ctx, "wenzhouyao_use_cnt").Int()
		if err != nil {
			fmt.Println("rds.Get fail, err:", err.Error())
		} else {
			fmt.Println("wenzhouyao_use_cnt:", cnt)
		}
	})
}

func TestRedisGet(t *testing.T) {
	ctx := context.Background()
	test(ctx, t, configs.REDIS_TYPE, func() {
		rds := db.RedisClient
		r, err := rds.Get(ctx, "wenzhouyao_use_cnt").Int()
		if err != nil {
			fmt.Println("rds.Incr Get, err:", err.Error())
		} else {
			fmt.Println("wenzhouyao_use_cnt =", r)
		}
	})
}
