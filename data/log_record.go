package data

type LogRecordType = byte

const (
	LogRecordNormal LogRecordType = iota
	LogRecordDeleted
)

type LogRecord struct {
	Key   []byte
	Value []byte
	Type  LogRecordType
}
type LogRecordPos struct {
	Fid    uint32
	Offset int64
	//文件id和偏移,数据存到文件的名字，数据在文件中的位置
	//首字母大写表示可被外部访问
}

func EncodeLogRecord(logRecord *LogRecord) ([]byte, int64) {
	return nil, 0
}
