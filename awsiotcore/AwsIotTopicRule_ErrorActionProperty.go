package awsiotcore


// Experimental.
type AwsIotTopicRule_ErrorActionProperty struct {
	// cloudwatch_alarm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#cloudwatch_alarm AwsIotTopicRule#cloudwatch_alarm}
	// Experimental.
	CloudwatchAlarm *AwsIotTopicRule_ErrorActionCloudwatchAlarmProperty `field:"optional" json:"cloudwatchAlarm" yaml:"cloudwatchAlarm"`
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#cloudwatch_logs AwsIotTopicRule#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *AwsIotTopicRule_ErrorActionCloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// cloudwatch_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#cloudwatch_metric AwsIotTopicRule#cloudwatch_metric}
	// Experimental.
	CloudwatchMetric *AwsIotTopicRule_ErrorActionCloudwatchMetricProperty `field:"optional" json:"cloudwatchMetric" yaml:"cloudwatchMetric"`
	// dynamodb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#dynamodb AwsIotTopicRule#dynamodb}
	// Experimental.
	Dynamodb *AwsIotTopicRule_ErrorActionDynamodbProperty `field:"optional" json:"dynamodb" yaml:"dynamodb"`
	// dynamodbv2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#dynamodbv2 AwsIotTopicRule#dynamodbv2}
	// Experimental.
	Dynamodbv2 *AwsIotTopicRule_ErrorActionDynamodbv2Property `field:"optional" json:"dynamodbv2" yaml:"dynamodbv2"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#elasticsearch AwsIotTopicRule#elasticsearch}
	// Experimental.
	Elasticsearch *AwsIotTopicRule_ErrorActionElasticsearchProperty `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#firehose AwsIotTopicRule#firehose}
	// Experimental.
	Firehose *AwsIotTopicRule_ErrorActionFirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#http AwsIotTopicRule#http}
	// Experimental.
	Http *AwsIotTopicRule_ErrorActionHttpProperty `field:"optional" json:"http" yaml:"http"`
	// iot_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#iot_analytics AwsIotTopicRule#iot_analytics}
	// Experimental.
	IotAnalytics *AwsIotTopicRule_ErrorActionIotAnalyticsProperty `field:"optional" json:"iotAnalytics" yaml:"iotAnalytics"`
	// iot_events block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#iot_events AwsIotTopicRule#iot_events}
	// Experimental.
	IotEvents *AwsIotTopicRule_ErrorActionIotEventsProperty `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// kafka block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#kafka AwsIotTopicRule#kafka}
	// Experimental.
	Kafka *AwsIotTopicRule_ErrorActionKafkaProperty `field:"optional" json:"kafka" yaml:"kafka"`
	// kinesis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#kinesis AwsIotTopicRule#kinesis}
	// Experimental.
	Kinesis *AwsIotTopicRule_ErrorActionKinesisProperty `field:"optional" json:"kinesis" yaml:"kinesis"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#lambda AwsIotTopicRule#lambda}
	// Experimental.
	Lambda *AwsIotTopicRule_ErrorActionLambdaProperty `field:"optional" json:"lambda" yaml:"lambda"`
	// republish block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#republish AwsIotTopicRule#republish}
	// Experimental.
	Republish *AwsIotTopicRule_ErrorActionRepublishProperty `field:"optional" json:"republish" yaml:"republish"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#s3 AwsIotTopicRule#s3}
	// Experimental.
	S3 *AwsIotTopicRule_ErrorActionS3Property `field:"optional" json:"s3" yaml:"s3"`
	// sns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#sns AwsIotTopicRule#sns}
	// Experimental.
	Sns *AwsIotTopicRule_ErrorActionSnsProperty `field:"optional" json:"sns" yaml:"sns"`
	// sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#sqs AwsIotTopicRule#sqs}
	// Experimental.
	Sqs *AwsIotTopicRule_ErrorActionSqsProperty `field:"optional" json:"sqs" yaml:"sqs"`
	// step_functions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#step_functions AwsIotTopicRule#step_functions}
	// Experimental.
	StepFunctions *AwsIotTopicRule_ErrorActionStepFunctionsProperty `field:"optional" json:"stepFunctions" yaml:"stepFunctions"`
	// timestream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#timestream AwsIotTopicRule#timestream}
	// Experimental.
	Timestream *AwsIotTopicRule_ErrorActionTimestreamProperty `field:"optional" json:"timestream" yaml:"timestream"`
}

