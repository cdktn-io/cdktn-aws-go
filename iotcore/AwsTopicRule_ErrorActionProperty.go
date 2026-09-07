package iotcore


// Experimental.
type AwsTopicRule_ErrorActionProperty struct {
	// cloudwatch_alarm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#cloudwatch_alarm AwsTopicRule#cloudwatch_alarm}
	// Experimental.
	CloudwatchAlarm *AwsTopicRule_ErrorActionCloudwatchAlarmProperty `field:"optional" json:"cloudwatchAlarm" yaml:"cloudwatchAlarm"`
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#cloudwatch_logs AwsTopicRule#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *AwsTopicRule_ErrorActionCloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// cloudwatch_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#cloudwatch_metric AwsTopicRule#cloudwatch_metric}
	// Experimental.
	CloudwatchMetric *AwsTopicRule_ErrorActionCloudwatchMetricProperty `field:"optional" json:"cloudwatchMetric" yaml:"cloudwatchMetric"`
	// dynamodb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#dynamodb AwsTopicRule#dynamodb}
	// Experimental.
	Dynamodb *AwsTopicRule_ErrorActionDynamodbProperty `field:"optional" json:"dynamodb" yaml:"dynamodb"`
	// dynamodbv2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#dynamodbv2 AwsTopicRule#dynamodbv2}
	// Experimental.
	Dynamodbv2 *AwsTopicRule_ErrorActionDynamodbv2Property `field:"optional" json:"dynamodbv2" yaml:"dynamodbv2"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#elasticsearch AwsTopicRule#elasticsearch}
	// Experimental.
	Elasticsearch *AwsTopicRule_ErrorActionElasticsearchProperty `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#firehose AwsTopicRule#firehose}
	// Experimental.
	Firehose *AwsTopicRule_ErrorActionFirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#http AwsTopicRule#http}
	// Experimental.
	Http *AwsTopicRule_ErrorActionHttpProperty `field:"optional" json:"http" yaml:"http"`
	// iot_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#iot_analytics AwsTopicRule#iot_analytics}
	// Experimental.
	IotAnalytics *AwsTopicRule_ErrorActionIotAnalyticsProperty `field:"optional" json:"iotAnalytics" yaml:"iotAnalytics"`
	// iot_events block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#iot_events AwsTopicRule#iot_events}
	// Experimental.
	IotEvents *AwsTopicRule_ErrorActionIotEventsProperty `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// kafka block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#kafka AwsTopicRule#kafka}
	// Experimental.
	Kafka *AwsTopicRule_ErrorActionKafkaProperty `field:"optional" json:"kafka" yaml:"kafka"`
	// kinesis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#kinesis AwsTopicRule#kinesis}
	// Experimental.
	Kinesis *AwsTopicRule_ErrorActionKinesisProperty `field:"optional" json:"kinesis" yaml:"kinesis"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#lambda AwsTopicRule#lambda}
	// Experimental.
	Lambda *AwsTopicRule_ErrorActionLambdaProperty `field:"optional" json:"lambda" yaml:"lambda"`
	// republish block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#republish AwsTopicRule#republish}
	// Experimental.
	Republish *AwsTopicRule_ErrorActionRepublishProperty `field:"optional" json:"republish" yaml:"republish"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#s3 AwsTopicRule#s3}
	// Experimental.
	S3 *AwsTopicRule_ErrorActionS3Property `field:"optional" json:"s3" yaml:"s3"`
	// sns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#sns AwsTopicRule#sns}
	// Experimental.
	Sns *AwsTopicRule_ErrorActionSnsProperty `field:"optional" json:"sns" yaml:"sns"`
	// sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#sqs AwsTopicRule#sqs}
	// Experimental.
	Sqs *AwsTopicRule_ErrorActionSqsProperty `field:"optional" json:"sqs" yaml:"sqs"`
	// step_functions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#step_functions AwsTopicRule#step_functions}
	// Experimental.
	StepFunctions *AwsTopicRule_ErrorActionStepFunctionsProperty `field:"optional" json:"stepFunctions" yaml:"stepFunctions"`
	// timestream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#timestream AwsTopicRule#timestream}
	// Experimental.
	Timestream *AwsTopicRule_ErrorActionTimestreamProperty `field:"optional" json:"timestream" yaml:"timestream"`
}

