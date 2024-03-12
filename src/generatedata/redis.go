package generatedata

import (
	"context"
	"log"
	"strconv"
	"time"

	"git.services.wait/chenwx/cwxgoweb/src/metrics"
	"git.services.wait/chenwx/cwxgoweb/src/unit"
	"github.com/redis/go-redis/v9"
)

var redisCmdRunNum int64
var ctx = context.Background()

func redisTask(cfg GenerateConf) {
	log.Println("stress: --- generate data redis start")
	// var count int64         // 命令执行次数
	var scanEndDelNum int64 // 批量删除时实际删除数量
	var delStatusNum int64  // 执行批量删除时的 count 数

	// conf, err := config.GetRedisConfig()
	// if err != nil {
	// 	log.Println("get redis conf error, exit")
	// 	return
	// }

	rdb := getRedisSession(cfg.Redis.Address, cfg.Redis.Port)
	defer rdb.Close()

	// 定时计数器
	go unit.CountNumTicker("redis", &redisCmdRunNum)

	for {
		setKey_str(rdb, "str:", 0, unit.RandNumInt64Length(90000))
		redisCmdRunNum += 1
		metrics.RedisCmdNum.Add(1)

		setKey_str(rdb, "ttlKey:", unit.RandTimeMinute(10), unit.RandNumInt64Length(50000))
		redisCmdRunNum += 1
		metrics.RedisCmdNum.Add(1)

		// 每 4 次操写入一个 hash 子键
		if redisCmdRunNum%4 == 0 {
			setHashData(rdb, "chenwx", unit.RandNumInt64Length(90000))
			redisCmdRunNum += 1
			metrics.RedisCmdNum.Inc()
		}

		// 每 3 次操写入一个 key
		if redisCmdRunNum%3 == 0 {
			setKey_str(rdb, "xxx20:", 0, unit.RandNumInt64Length(50000))
			redisCmdRunNum += 1
			metrics.RedisCmdNum.Inc()
		}
		// 每 10 次操写入一个 hash 子键
		if redisCmdRunNum%10 == 0 {
			setKey_str(rdb, "xxx30:", 0, unit.RandNumInt64Length(90000))
			setHashData(rdb, "chenjia", unit.RandNumInt64Length(50000))
			redisCmdRunNum += 2
			metrics.RedisCmdNum.Add(2)
		}

		// 大概执行 10000 次命令, 就进行一次批量删除
		if redisCmdRunNum%10000 <= 1000 {
			if redisCmdRunNum-delStatusNum > 1000 {
				scanEndDelNum = delScanStrKey(rdb, 2000, "str:*")
				redisCmdRunNum += scanEndDelNum
				metrics.RedisCmdNum.Add(float64(scanEndDelNum))
				scanEndDelNum = delScanStrKey(rdb, 1000, "xxx20:*")
				redisCmdRunNum += scanEndDelNum
				metrics.RedisCmdNum.Add(float64(scanEndDelNum))
				scanEndDelNum = delScanStrKey(rdb, 1000, "xxx30:*")
				redisCmdRunNum += scanEndDelNum
				metrics.RedisCmdNum.Add(float64(scanEndDelNum))

				// log.Println("redis: del hash key chenwx, chenjia")
				redisCmdRunNum += 1
				rdb.Del(ctx, "chenwx")
				metrics.RedisCmdNum.Inc()
				redisCmdRunNum += 1
				rdb.Del(ctx, "chenjia")
				delStatusNum = redisCmdRunNum
				metrics.ScanDelRedisNum.Inc()
				metrics.RedisCmdNum.Inc()
			}

		}

		int_sleep, _ := strconv.Atoi(cfg.Redis.SleepMs)
		// 每轮操作中的延迟暂停
		time.Sleep(unit.RandTimeMillisecond(int_sleep))
	}

}

func getRedisSession(ip string, port string) *redis.Client {
	cacheAddr := ip + ":" + port
	client := redis.NewClient(&redis.Options{
		Addr:     cacheAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping(ctx).Result()

	if err != nil {
		log.Println("generatedata: link redis server error")
		panic(err)
	}

	return client
}

// func getListkey(r *redis.Client, list string) (string, error) {
// 	cmd := r.LPop(ctx, list)
// 	// cmd := r.BLPop(ctx, 2, list)
// 	data, err := cmd.Result()
// 	if err != nil {
// 		return "", err
// 	}
// 	return data, nil
// }

// func bGetListkey(r *redis.Client, list string, timeout time.Duration) ([]string, error) {
// 	cmd := r.BLPop(ctx, timeout, list)
// 	data, err := cmd.Result()
// 	if err != nil {
// 		return []string{}, err
// 	}
// 	return data, nil
// }

// func showKeys(r *redis.Client, keys []string) {
// 	for _, key := range keys {
// 		sType, err := r.Type(ctx, key).Result()
// 		if err != nil {
// 			log.Println("generatedata - redis: get type failed :", err)
// 			return
// 		}

// 		if sType == "string" {
// 			val, err := r.Get(ctx, key).Result()
// 			if err != nil {
// 				log.Println("generatedata - redis: get key values failed err:", err)
// 				return
// 			}
// 			log.Printf("key :%v ,value :%v\n", key, val)
// 		} else if sType == "list" {
// 			val, err := r.LPop(ctx, key).Result()
// 			if err != nil {
// 				log.Println("generatedata - redis: get list value failed :", err)
// 				return
// 			}
// 			log.Printf("generatedata - redis: key:%v value:%v\n", key, val)
// 		}
// 	}
// }

// 扫描删除 key, 返回实际删除数量
func delScanStrKey(r *redis.Client, number int64, header string) int64 {

	// number 传入的期望删除的 key 数量

	var endNum = number        // 剩余期望删除 key 数量
	var delNum int64 = 0       // 实际删除 key 数量
	var oneScanNum int64 = 500 // 单次扫描数量
	var oneScanKeyNum int64    // 单次扫描实际扫出来的 key 数量
	var scanNum int64 = 0      // 实际扫描次数
	var next_cursor uint64 = 0 // 下次扫描的游标位置

	// 进入循环删除
	for endNum > 0 {
		data := r.Scan(ctx, next_cursor, header, oneScanNum)
		scanNum += 1
		keys, cursor, err := data.Result()
		next_cursor = cursor
		if err != nil {
			log.Println("generatedata - redis: scan keys failed err:", err)
			break
		}

		oneScanKeyNum = int64(len(keys))
		// log.Println("redis: scan oneScanKeyNum:", oneScanKeyNum)

		// 游标为 0
		if cursor == 0 {
			// keys 内没有数据, 要么是第一次小集合扫, 要么是最后一次扫
			if oneScanKeyNum == 0 {
				log.Printf("generatedata - redis: scan %s cursor == 0, exit scan\n", header)
				break
			}

			// keys 内还有数据, 需要做最后一次处理, 再返回
			delKeys(r, keys)
			endNum -= oneScanKeyNum
			delNum += oneScanKeyNum
			break
		}

		// 如果单次扫描出来的 keys 多于 期望扫描的数量, 只删除一部分 key, 然后退出
		if oneScanKeyNum > endNum {
			newKeys := keys[0:endNum]
			delKeys(r, newKeys)
			delNum += endNum
			break
		}

		// 常规的普通扫描
		delKeys(r, keys)
		endNum -= oneScanKeyNum
		delNum += oneScanKeyNum
	}

	// log.Printf("redis: scan: %s number: %d, end scan num: %d, end del: %d\n", header, number, scanNum, delNum)

	return delNum

}

// 批量删除 key
func delKeys(r *redis.Client, keys []string) {
	for _, key := range keys {
		r.Del(ctx, key)
	}
}

// 为 hash key 设置一个 子 key
func setHashData(r *redis.Client, hashName string, randomNum int64) {
	hash_key := "m" + strconv.Itoa(int(randomNum))
	r.HSet(ctx, hashName, hash_key, randomNum)
}

// 设置一个 key
func setKey_str(r *redis.Client, header string, ttlTime time.Duration, randomNum int64) {
	key := header + strconv.Itoa(int(randomNum))
	r.Set(ctx, key, randomNum, ttlTime)
}
