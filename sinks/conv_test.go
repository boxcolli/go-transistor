package sinks

import (
	"testing"
	"time"

	pb "github.com/boxcolli/pepperlink/idl/gen/hello/v1"
	"github.com/stretchr/testify/assert"
)

func TestSqlMessageConverter(t *testing.T) {
	// 개발자 정의
	var mtow MTOW
	{
		mtow = MTOW{
			"hello": func([]byte, interface{}, time.Time) (string, error) {
				helloMsg := msg.(pb.Hello)
				return helloMsg.String(), nil
			},
		}
	}

	conv := NewSQLMessageConverter(mtow)
	{

	}

	// 컨버터 안에서 처리되는 방법
	var query string
	{
		msg := pb.Hello{
			Name: "Kim",
		}
		t := time.Now()
		query, _ = conv.GetWriteQuery("hello", []byte{1}, msg, t)
	}

	q := `INSERT INTO hello (id, name) VALUES (1, Kim)`
	assert.Equal(t, query, q)	// Test

	t.Log(query)
}
