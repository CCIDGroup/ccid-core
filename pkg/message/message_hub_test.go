package message

import (
	"github.com/streadway/amqp"
	"testing"
)

func TestRabbitMQ_ReadMQ(t *testing.T) {
	type fields struct {
		conn *amqp.Connection
		ch   *amqp.Channel
		q    amqp.Queue
	}
	type args struct {
		exID    string
		runID   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"",
			fields{

			},
			args {
				"testEXID",
				"testRUNID",
			},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rmq := (&RabbitMQ{}).InitAMQP(tt.args.runID)
			rmq.ReadMQ()
		})
	}
}