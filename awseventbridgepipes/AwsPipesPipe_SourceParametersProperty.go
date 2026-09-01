package awseventbridgepipes


// Experimental.
type AwsPipesPipe_SourceParametersProperty struct {
	// activemq_broker_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#activemq_broker_parameters AwsPipesPipe#activemq_broker_parameters}
	// Experimental.
	ActivemqBrokerParameters *AwsPipesPipe_ActivemqBrokerParametersProperty `field:"optional" json:"activemqBrokerParameters" yaml:"activemqBrokerParameters"`
	// dynamodb_stream_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#dynamodb_stream_parameters AwsPipesPipe#dynamodb_stream_parameters}
	// Experimental.
	DynamodbStreamParameters *AwsPipesPipe_DynamodbStreamParametersProperty `field:"optional" json:"dynamodbStreamParameters" yaml:"dynamodbStreamParameters"`
	// filter_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#filter_criteria AwsPipesPipe#filter_criteria}
	// Experimental.
	FilterCriteria *AwsPipesPipe_FilterCriteriaProperty `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// kinesis_stream_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#kinesis_stream_parameters AwsPipesPipe#kinesis_stream_parameters}
	// Experimental.
	KinesisStreamParameters *AwsPipesPipe_SourceParametersKinesisStreamParametersProperty `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// managed_streaming_kafka_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#managed_streaming_kafka_parameters AwsPipesPipe#managed_streaming_kafka_parameters}
	// Experimental.
	ManagedStreamingKafkaParameters *AwsPipesPipe_ManagedStreamingKafkaParametersProperty `field:"optional" json:"managedStreamingKafkaParameters" yaml:"managedStreamingKafkaParameters"`
	// rabbitmq_broker_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#rabbitmq_broker_parameters AwsPipesPipe#rabbitmq_broker_parameters}
	// Experimental.
	RabbitmqBrokerParameters *AwsPipesPipe_RabbitmqBrokerParametersProperty `field:"optional" json:"rabbitmqBrokerParameters" yaml:"rabbitmqBrokerParameters"`
	// self_managed_kafka_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#self_managed_kafka_parameters AwsPipesPipe#self_managed_kafka_parameters}
	// Experimental.
	SelfManagedKafkaParameters *AwsPipesPipe_SelfManagedKafkaParametersProperty `field:"optional" json:"selfManagedKafkaParameters" yaml:"selfManagedKafkaParameters"`
	// sqs_queue_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#sqs_queue_parameters AwsPipesPipe#sqs_queue_parameters}
	// Experimental.
	SqsQueueParameters *AwsPipesPipe_SourceParametersSqsQueueParametersProperty `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
}

