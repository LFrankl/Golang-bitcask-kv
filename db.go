package Golang_bitcask_kv

import (
	"go-bitcask/data"
	"sync"
)

// DB 存储引擎实例
type DB struct {
	mu         sync.RWMutex
	activeFile *data.DataFile            //当前活跃数据文件，可用于写入
	olderFiles map[uint32]*data.DataFile //旧的数据文件，只能用于读
}

func (db *DB) Put(key []byte, value []byte) error {
	if len(key) == 0 {
		//判断key是否有效
		return ErrKeyIsEmpty
	}
	//构造LogRecord结构体
	log_record := data.LogRecord{
		Key:   key,
		Value: value,
		Type:  data.LogRecordNormal,
	}
}

func (db *DB) appendLogRecord(logRecord *data.LogRecord) (*data.LogRecordPos, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	//判断当前活跃数据文件是否存在，db在没有写入的时候是没有文件生成的
	//如果为空则初始化数据文件
	if db.activeFile != nil {

	}
}

// 设置当前活跃文件
func (db *DB) setActiveFile() error {
	var initialFileId uint32 = 0
	if db.activeFile != nil {

	}
}
