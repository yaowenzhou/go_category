package pkg

import (
	"sync"
	"time"
)

const (
	machineBits  = int64(5)  //机器id位数
	serviceBits  = int64(5)  //服务id位数
	sequenceBits = int64(12) //序列id位数

	maxMachineID  = int64(-1) ^ (int64(-1) << machineBits)  //最大机器id
	maxServiceID  = int64(-1) ^ (int64(-1) << serviceBits)  //最大服务id
	maxSequenceID = int64(-1) ^ (int64(-1) << sequenceBits) //最大序列id

	timeLeft    = uint8(22) //时间id向左移位的量
	machineLeft = uint8(17) //机器id向左移位的量
	serviceLeft = uint8(12) //服务id向左移位的量

	twepoch = int64(1667972427000) //初始毫秒,时间是: Wed Nov  9 13:40:27 CST 2022
)

type Worker struct {
	sync.Mutex
	lastStamp  int64
	machineID  int64 //机器id,0~31
	serviceID  int64 //服务id,0~31
	sequenceID int64
}

func NewWorker(machineID, serviceID int64) *Worker {
	return &Worker{
		lastStamp:  0,
		machineID:  machineID,
		serviceID:  serviceID,
		sequenceID: 0,
	}
}

// 分布式系统中，NewWorker的参数需要动态获取
var work = NewWorker(30, 30)

// GetID 计算雪花id
func (w *Worker) GetID() int64 {
	w.Lock() //多线程互斥
	defer w.Unlock()
	mill := time.Now().UnixMilli()
	if mill == w.lastStamp {
		w.sequenceID = (w.sequenceID + 1) & maxSequenceID
		//当一个毫秒内分配的id数>4096个时，只能等待到下一毫秒去分配。
		if w.sequenceID == 0 {
			for mill > w.lastStamp {
				mill = time.Now().UnixMilli()
			}
		}
	} else {
		w.sequenceID = 0
	}

	w.lastStamp = mill
	id := (w.lastStamp-twepoch)<<timeLeft | w.machineID<<machineLeft | w.serviceID<<serviceLeft | w.sequenceID
	return id
}

// GetSnowFlake 获取雪花id
func GetSnowFlake() int64 {
	return work.GetID()
}
