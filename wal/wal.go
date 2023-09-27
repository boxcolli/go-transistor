package wal

import "time"

/*
문제:
	받은 메시지를 즉시 데이터베이스에 넣으면 이상적이지만 그 전에 노드가 강제종료되면 받은 메시지를 다 잃어버린다.

WAL:
	메시지를 받은 즉시 로컬 저장소에 큐 형태로 저장한다.
	따로 돌아가는 쓰레드가 메시지를 순서대로 처리한다.
	들어온 순서대로 할지, 타임스탬프 순서대로 할 지 아직 모르겠음. 키를 무엇으로 잡냐에 따라 달라짐
	주기적으로 큐를 비운다.
*/

type WAL interface {
	/*
		인터페이스:
			새 메시지를 쓰는 함수
			주기적으로 메시지를 비울 때 호출할 함수
	*/
	Push(topic string, topicId []byte, msg []byte, timestamp time.Time) error
	Flush(timeCursor time.Time) error	// 특정 timestamp 기준으로 전부 데이터베이스 삽입에 성공했다고 가정하고, 그것보다 오래된 메시지를 다 버린다.
}
