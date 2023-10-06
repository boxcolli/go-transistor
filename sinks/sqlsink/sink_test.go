package sqlsink

import (
	"database/sql"
	"testing"
	"time"

	pbhello "github.com/boxcolli/go-transistor/idl/gen/hello/v1"
	pb "github.com/boxcolli/go-transistor/idl/gen/link/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var topic = "hello"
var topicId = uuid.New().NodeID()
var dataId = uuid.New().NodeID()
var dataName = "Alice"

// { Create, Update, Delete }
var datas = []pbhello.Hello{
	{
		Id: wrapperspb.Bytes(dataId),
		Name: wrapperspb.String(dataName),
		Age: wrapperspb.Int32(20),
	},
	{
		Id: wrapperspb.Bytes(dataId),
		Age: wrapperspb.Int32(30),
	},
	{
		Id: wrapperspb.Bytes(dataId),
	},
}

// { Create, Update, Delete }
var testset = []pb.PublishRequest{
	{
		Topic: wrapperspb.String(topic),
		TopicId: wrapperspb.Bytes(topicId),
		Method: pb.Method_METHOD_CREATE,
		Data: nil,
		Timestamp: timestamppb.Now(),
	},
	{
		Topic: wrapperspb.String(topic),
		TopicId: wrapperspb.Bytes(topicId),
		Method: pb.Method_METHOD_UPDATE,
		Data: nil,
		Timestamp: timestamppb.Now(),
	},
	{
		Topic: wrapperspb.String(topic),
		TopicId: wrapperspb.Bytes(topicId),
		Method: pb.Method_METHOD_DELETE,
		Data: nil,
		Timestamp: timestamppb.Now(),
	},
}
var Q = new(Query)

func prepareQuery() {
	{
		var err error
		db, err = sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/test?parseTime=true")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		log.Println("Connected!")

		Q.DB = db
	}
	{
		{
			stmt, err := Q.DB.Prepare(`
			INSERT INTO hello (id, name, name_tp, age, age_tp) VALUES (?, ?, ?, ?, ?)
			`)
			assert.NoError(t, err)
			Q.CS[topic] = stmt
		}
		{
			stmt, err := Q.DB.Prepare(`
			INSERT INTO hello (id, name, name_tp, age, age_tp)
			VALUES (?, ?, ?, ?, ?)
			ON DUPLICATE KEY UPDATE
			name = IF(name_tp > VALUES(name_tp), name, VALUES(name)),
			name_tp = IF(name_tp > VALUES(name_tp), name_tp, VALUES(name_tp)),
			timestamp_column = IF(timestamp_column > VALUES(timestamp_column), timestamp_column, VALUES(timestamp_column));
			`)
			assert.NoError(t, err)
			Q.CS[topic] = stmt
		}
	}
	{
		
		Q.Create["hello"] = func([]byte, interface{}, time.Time) error {
			query := `
			INSERT INTO hello (id, name, name_tp, age, age_tp) VALUES (?, ?, ?, ?, ?)
			`
			Q.db.
			return nil
		}
		Q.Update["hello"] = nil
		Q.Delete["hello"] = nil
	}
}

func TestSQLSink(t *testing.T) {
	prepareTestset()

	{
		// Grab function
		f := mtow.Create["hello"]
		assert.NotNil(t, f)

		// Grab Data
		msg := &testset[0]
		hello := new(pbhello.Hello)
		assert.NoError(t, testset[0].GetData().UnmarshalTo(hello))
		
		// Convert
		query, err :=f(msg.TopicId.GetValue(), hello, msg.Timestamp.AsTime())
		assert.NoError(t, err)

		// Print
		t.Log("CREATE:", query)
	}

	{
		// Grab function
		f := mtow.Create["hello"]
		assert.NotNil(t, f)

		// Grab Data
		msg := &testset[1]
		hello := new(pbhello.Hello)
		assert.NoError(t, testset[0].GetData().UnmarshalTo(hello))
		
		// Convert
		query, err :=f(msg.TopicId.GetValue(), hello, msg.Timestamp.AsTime())
		assert.NoError(t, err)

		// Print
		t.Log("UPDATE:", query)
	}

	{
		// Grab function
		f := mtow.Create["hello"]
		assert.NotNil(t, f)

		// Grab Data
		msg := &testset[2]
		hello := new(pbhello.Hello)
		assert.NoError(t, testset[0].GetData().UnmarshalTo(hello))
		
		// Convert
		query, err :=f(msg.TopicId.GetValue(), hello, msg.Timestamp.AsTime())
		assert.NoError(t, err)

		// Print
		t.Log("DELETE:", query)
	}
}

func prepareTestset() {
	// Prepare test set
	{
		// Put data[i] -> testset[i].Data
		for i := range testset {
			any, err := anypb.New(&datas[i])
			if err != nil {
				t.Fatal(err)
			}
			testset[i].Data = any
		}
	}
}