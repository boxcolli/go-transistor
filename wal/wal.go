package wal

import "time"

type WAL interface {
	/*
		인터페이스:
			새 메시지를 쓰는 함수
			주기적으로 메시지를 비울 때 호출할 함수
	*/
	Push(topic string, topicId []byte, msg []byte, timestamp time.Time) error
	Flush(timeCursor time.Time) error	// 특정 timestamp 기준으로 전부 데이터베이스 삽입에 성공했다고 가정하고, 그것보다 오래된 메시지를 다 버린다.
}
