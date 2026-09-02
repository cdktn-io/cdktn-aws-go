package awsiotcore


// Experimental.
type TfTopicRule_ErrorActionProperty struct {
	// cloudwatch_alarm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#cloudwatch_alarm TfTopicRule#cloudwatch_alarm}
	// Experimental.
	CloudwatchAlarm *TfTopicRule_ErrorActionCloudwatchAlarmProperty `field:"optional" json:"cloudwatchAlarm" yaml:"cloudwatchAlarm"`
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#cloudwatch_logs TfTopicRule#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *TfTopicRule_ErrorActionCloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// cloudwatch_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#cloudwatch_metric TfTopicRule#cloudwatch_metric}
	// Experimental.
	CloudwatchMetric *TfTopicRule_ErrorActionCloudwatchMetricProperty `field:"optional" json:"cloudwatchMetric" yaml:"cloudwatchMetric"`
	// dynamodb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#dynamodb TfTopicRule#dynamodb}
	// Experimental.
	Dynamodb *TfTopicRule_ErrorActionDynamodbProperty `field:"optional" json:"dynamodb" yaml:"dynamodb"`
	// dynamodbv2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#dynamodbv2 TfTopicRule#dynamodbv2}
	// Experimental.
	Dynamodbv2 *TfTopicRule_ErrorActionDynamodbv2Property `field:"optional" json:"dynamodbv2" yaml:"dynamodbv2"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#elasticsearch TfTopicRule#elasticsearch}
	// Experimental.
	Elasticsearch *TfTopicRule_ErrorActionElasticsearchProperty `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#firehose TfTopicRule#firehose}
	// Experimental.
	Firehose *TfTopicRule_ErrorActionFirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#http TfTopicRule#http}
	// Experimental.
	Http *TfTopicRule_ErrorActionHttpProperty `field:"optional" json:"http" yaml:"http"`
	// iot_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#iot_analytics TfTopicRule#iot_analytics}
	// Experimental.
	IotAnalytics *TfTopicRule_ErrorActionIotAnalyticsProperty `field:"optional" json:"iotAnalytics" yaml:"iotAnalytics"`
	// iot_events block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#iot_events TfTopicRule#iot_events}
	// Experimental.
	IotEvents *TfTopicRule_ErrorActionIotEventsProperty `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// kafka block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#kafka TfTopicRule#kafka}
	// Experimental.
	Kafka *TfTopicRule_ErrorActionKafkaProperty `field:"optional" json:"kafka" yaml:"kafka"`
	// kinesis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#kinesis TfTopicRule#kinesis}
	// Experimental.
	Kinesis *TfTopicRule_ErrorActionKinesisProperty `field:"optional" json:"kinesis" yaml:"kinesis"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#lambda TfTopicRule#lambda}
	// Experimental.
	Lambda *TfTopicRule_ErrorActionLambdaProperty `field:"optional" json:"lambda" yaml:"lambda"`
	// republish block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#republish TfTopicRule#republish}
	// Experimental.
	Republish *TfTopicRule_ErrorActionRepublishProperty `field:"optional" json:"republish" yaml:"republish"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#s3 TfTopicRule#s3}
	// Experimental.
	S3 *TfTopicRule_ErrorActionS3Property `field:"optional" json:"s3" yaml:"s3"`
	// sns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#sns TfTopicRule#sns}
	// Experimental.
	Sns *TfTopicRule_ErrorActionSnsProperty `field:"optional" json:"sns" yaml:"sns"`
	// sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#sqs TfTopicRule#sqs}
	// Experimental.
	Sqs *TfTopicRule_ErrorActionSqsProperty `field:"optional" json:"sqs" yaml:"sqs"`
	// step_functions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#step_functions TfTopicRule#step_functions}
	// Experimental.
	StepFunctions *TfTopicRule_ErrorActionStepFunctionsProperty `field:"optional" json:"stepFunctions" yaml:"stepFunctions"`
	// timestream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#timestream TfTopicRule#timestream}
	// Experimental.
	Timestream *TfTopicRule_ErrorActionTimestreamProperty `field:"optional" json:"timestream" yaml:"timestream"`
}

