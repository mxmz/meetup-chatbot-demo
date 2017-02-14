package aws

import (
	"fmt"
	"log"
	"math/rand"
	"telegram-meetup/env"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main2() {
	env := env.LoadEnv()
	env.Export()

	sess := session.New(&aws.Config{
		Region:     aws.String("us-west-2"),
		DisableSSL: aws.Bool(true),
	})

	item := map[string]*dynamodb.AttributeValue{
		"userid": {
			S: aws.String("a3"),
		},
		"pippoa": {
			N: aws.String("424"),
		},
	}

	params := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("users"), // Required
	}
	svc := dynamodb.New(sess)
	_, err := svc.PutItem(params)
	log.Println(err)

	params2 := &dynamodb.UpdateItemInput{
		TableName: aws.String("meetups_users"), // Required
		Key: map[string]*dynamodb.AttributeValue{
			"meetup_user": {S: aws.String("1114")},
		},
		ExpressionAttributeNames: map[string]*string{
			"#c": aws.String("meetup"),
			"#t": aws.String("tag"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":c": {S: aws.String("PIPPOAA")},
			":t": {S: aws.String("eoritueori")},
		},
		UpdateExpression: aws.String("SET #c = :c, #t = :t"),
		ReturnValues:     aws.String("ALL_NEW"),
	}

	r, err := svc.UpdateItem(params2)
	log.Println(err, r)

	/*
		params := &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]*dynamodb.WriteRequest{
				"users": {
					&dynamodb.WriteRequest{
						PutRequest: &dynamodb.PutRequest{
							Item: map[string]*dynamodb.AttributeValue{
								"userid": {
									S: aws.String("a1"),
								},
								"pippoa": {
									N: aws.String("424"),
								},
							},
						},
					},
				},
			},
		}

		svc := dynamodb.New(sess)
		_, err := svc.BatchWriteItem(params)
		log.Println(err)
	*/
	_ = err
	_ = svc
	_ = params
	_ = sess
}

func main1() {
	env := env.LoadEnv()
	env.Export()

	sess := session.New(&aws.Config{
		Region:     aws.String("us-west-2"),
		DisableSSL: aws.Bool(true),
	})

	svc := sqs.New(sess)

	queues, err := svc.ListQueues(nil)
	if err != nil {
		panic(err)
	}

	url := queues.QueueUrls[0]

	log.Println(*url)
	resp, err := svc.SendMessage(&sqs.SendMessageInput{
		QueueUrl:               url,
		MessageBody:            aws.String("bodaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaysa"),
		MessageGroupId:         aws.String("patagora"),
		MessageDeduplicationId: aws.String("patagorssssssssssssssssssssaassas"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("MD5 of body:", *resp.MD5OfMessageBody)
	receive_params := &sqs.ReceiveMessageInput{
		QueueUrl:            url,
		MaxNumberOfMessages: aws.Int64(3),
		VisibilityTimeout:   aws.Int64(30),
		WaitTimeSeconds:     aws.Int64(20),
	}
	receive_resp, err := svc.ReceiveMessage(receive_params)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("[Receive message] \n%v \n\n", receive_resp)
	return
	// Delete message
	for _, message := range receive_resp.Messages {
		delete_params := &sqs.DeleteMessageInput{
			QueueUrl:      url,
			ReceiptHandle: message.ReceiptHandle, // Required

		}
		_, err := svc.DeleteMessage(delete_params) // No response returned when successed.
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("[Delete message] \nMessage ID: %s has beed deleted.\n\n", *message.MessageId)
	}

}

func randStr(c int) string {
	var s = ""
	for n := 0; n < c; n++ {
		var ch = rune(rand.Intn(26) + 65)
		s = s + string(ch)
	}
	return s
}
