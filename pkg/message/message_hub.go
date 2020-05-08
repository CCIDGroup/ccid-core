/*
 * Copyright 2020 The CCID Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the 'License');
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http: //www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an 'AS IS' BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package message

import (
	"fmt"
	"github.com/CCIDGroup/ccid-core/utils"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}


func (rmq *RabbitMQ) InitAMQP(runID string) *RabbitMQ{

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		utils.LogError(err,"Failed to rmq.connect to RabbitMQ")
	}
	rmq.conn = conn
	rmq.ch, err = rmq.conn.Channel()
	if err != nil {
		utils.LogError(err, "Failed to open a channel")
	}
	rmq.declareQueue(runID)
	return rmq

}


func (rmq *RabbitMQ) declareQueue(queueID string){
	var err error
	rmq.q, err = rmq.ch.QueueDeclare(
		queueID,    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		utils.LogError(err, "Failed to declare a queue")
	}

}


func (rmq *RabbitMQ) WriteToMQ(ch *chan string) {
	for {
		val, ok := <- *ch
		if ok == false {
			break
		} else {
			fmt.Print(val)
			err := rmq.ch.Publish(
				"",         // exchange
				rmq.q.Name, // routing key
				false, // mandatory
				false, // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(val),
				})
			if err != nil {
				utils.LogError(err, "Failed to bind a queue")
			}
		}
	}
}

func (rmq *RabbitMQ) ReadMQ() *chan string {



	msgs, err := rmq.ch.Consume(
		rmq.q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		utils.LogError(err, "Failed to register a consumer")
	}


	result := make(chan string)

	go func() {
		for d := range msgs {
			result <-  string(d.Body)
		}
	}()

	utils.LogMsg(" [*] Waiting for logs. To exit press CTRL+C")
	return &result
}



