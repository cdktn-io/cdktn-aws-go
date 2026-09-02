package awseventbridgepipes


// Experimental.
type TfPipe_SourceParametersProperty struct {
	// activemq_broker_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#activemq_broker_parameters TfPipe#activemq_broker_parameters}
	// Experimental.
	ActivemqBrokerParameters *TfPipe_ActivemqBrokerParametersProperty `field:"optional" json:"activemqBrokerParameters" yaml:"activemqBrokerParameters"`
	// dynamodb_stream_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#dynamodb_stream_parameters TfPipe#dynamodb_stream_parameters}
	// Experimental.
	DynamodbStreamParameters *TfPipe_DynamodbStreamParametersProperty `field:"optional" json:"dynamodbStreamParameters" yaml:"dynamodbStreamParameters"`
	// filter_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#filter_criteria TfPipe#filter_criteria}
	// Experimental.
	FilterCriteria *TfPipe_FilterCriteriaProperty `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// kinesis_stream_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#kinesis_stream_parameters TfPipe#kinesis_stream_parameters}
	// Experimental.
	KinesisStreamParameters *TfPipe_SourceParametersKinesisStreamParametersProperty `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// managed_streaming_kafka_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#managed_streaming_kafka_parameters TfPipe#managed_streaming_kafka_parameters}
	// Experimental.
	ManagedStreamingKafkaParameters *TfPipe_ManagedStreamingKafkaParametersProperty `field:"optional" json:"managedStreamingKafkaParameters" yaml:"managedStreamingKafkaParameters"`
	// rabbitmq_broker_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#rabbitmq_broker_parameters TfPipe#rabbitmq_broker_parameters}
	// Experimental.
	RabbitmqBrokerParameters *TfPipe_RabbitmqBrokerParametersProperty `field:"optional" json:"rabbitmqBrokerParameters" yaml:"rabbitmqBrokerParameters"`
	// self_managed_kafka_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#self_managed_kafka_parameters TfPipe#self_managed_kafka_parameters}
	// Experimental.
	SelfManagedKafkaParameters *TfPipe_SelfManagedKafkaParametersProperty `field:"optional" json:"selfManagedKafkaParameters" yaml:"selfManagedKafkaParameters"`
	// sqs_queue_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#sqs_queue_parameters TfPipe#sqs_queue_parameters}
	// Experimental.
	SqsQueueParameters *TfPipe_SourceParametersSqsQueueParametersProperty `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
}

