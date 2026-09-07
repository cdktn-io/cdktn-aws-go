package eventbridgepipes


// Experimental.
type AwsPipe_SourceParametersProperty struct {
	// activemq_broker_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#activemq_broker_parameters AwsPipe#activemq_broker_parameters}
	// Experimental.
	ActivemqBrokerParameters *AwsPipe_ActivemqBrokerParametersProperty `field:"optional" json:"activemqBrokerParameters" yaml:"activemqBrokerParameters"`
	// dynamodb_stream_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#dynamodb_stream_parameters AwsPipe#dynamodb_stream_parameters}
	// Experimental.
	DynamodbStreamParameters *AwsPipe_DynamodbStreamParametersProperty `field:"optional" json:"dynamodbStreamParameters" yaml:"dynamodbStreamParameters"`
	// filter_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#filter_criteria AwsPipe#filter_criteria}
	// Experimental.
	FilterCriteria *AwsPipe_FilterCriteriaProperty `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// kinesis_stream_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#kinesis_stream_parameters AwsPipe#kinesis_stream_parameters}
	// Experimental.
	KinesisStreamParameters *AwsPipe_SourceParametersKinesisStreamParametersProperty `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// managed_streaming_kafka_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#managed_streaming_kafka_parameters AwsPipe#managed_streaming_kafka_parameters}
	// Experimental.
	ManagedStreamingKafkaParameters *AwsPipe_ManagedStreamingKafkaParametersProperty `field:"optional" json:"managedStreamingKafkaParameters" yaml:"managedStreamingKafkaParameters"`
	// rabbitmq_broker_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#rabbitmq_broker_parameters AwsPipe#rabbitmq_broker_parameters}
	// Experimental.
	RabbitmqBrokerParameters *AwsPipe_RabbitmqBrokerParametersProperty `field:"optional" json:"rabbitmqBrokerParameters" yaml:"rabbitmqBrokerParameters"`
	// self_managed_kafka_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#self_managed_kafka_parameters AwsPipe#self_managed_kafka_parameters}
	// Experimental.
	SelfManagedKafkaParameters *AwsPipe_SelfManagedKafkaParametersProperty `field:"optional" json:"selfManagedKafkaParameters" yaml:"selfManagedKafkaParameters"`
	// sqs_queue_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#sqs_queue_parameters AwsPipe#sqs_queue_parameters}
	// Experimental.
	SqsQueueParameters *AwsPipe_SourceParametersSqsQueueParametersProperty `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
}

