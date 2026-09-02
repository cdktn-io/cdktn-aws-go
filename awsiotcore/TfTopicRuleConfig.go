package awsiotcore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTopicRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#enabled TfTopicRule#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#name TfTopicRule#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#sql TfTopicRule#sql}.
	// Experimental.
	Sql *string `field:"required" json:"sql" yaml:"sql"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#sql_version TfTopicRule#sql_version}.
	// Experimental.
	SqlVersion *string `field:"required" json:"sqlVersion" yaml:"sqlVersion"`
	// cloudwatch_alarm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#cloudwatch_alarm TfTopicRule#cloudwatch_alarm}
	// Experimental.
	CloudwatchAlarm interface{} `field:"optional" json:"cloudwatchAlarm" yaml:"cloudwatchAlarm"`
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#cloudwatch_logs TfTopicRule#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs interface{} `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// cloudwatch_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#cloudwatch_metric TfTopicRule#cloudwatch_metric}
	// Experimental.
	CloudwatchMetric interface{} `field:"optional" json:"cloudwatchMetric" yaml:"cloudwatchMetric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#description TfTopicRule#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dynamodb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#dynamodb TfTopicRule#dynamodb}
	// Experimental.
	Dynamodb interface{} `field:"optional" json:"dynamodb" yaml:"dynamodb"`
	// dynamodbv2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#dynamodbv2 TfTopicRule#dynamodbv2}
	// Experimental.
	Dynamodbv2 interface{} `field:"optional" json:"dynamodbv2" yaml:"dynamodbv2"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#elasticsearch TfTopicRule#elasticsearch}
	// Experimental.
	Elasticsearch interface{} `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// error_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#error_action TfTopicRule#error_action}
	// Experimental.
	ErrorAction *TfTopicRule_ErrorActionProperty `field:"optional" json:"errorAction" yaml:"errorAction"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#firehose TfTopicRule#firehose}
	// Experimental.
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#http TfTopicRule#http}
	// Experimental.
	Http interface{} `field:"optional" json:"http" yaml:"http"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#id TfTopicRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// iot_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#iot_analytics TfTopicRule#iot_analytics}
	// Experimental.
	IotAnalytics interface{} `field:"optional" json:"iotAnalytics" yaml:"iotAnalytics"`
	// iot_events block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#iot_events TfTopicRule#iot_events}
	// Experimental.
	IotEvents interface{} `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// kafka block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#kafka TfTopicRule#kafka}
	// Experimental.
	Kafka interface{} `field:"optional" json:"kafka" yaml:"kafka"`
	// kinesis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#kinesis TfTopicRule#kinesis}
	// Experimental.
	Kinesis interface{} `field:"optional" json:"kinesis" yaml:"kinesis"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#lambda TfTopicRule#lambda}
	// Experimental.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#region TfTopicRule#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// republish block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#republish TfTopicRule#republish}
	// Experimental.
	Republish interface{} `field:"optional" json:"republish" yaml:"republish"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#s3 TfTopicRule#s3}
	// Experimental.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// sns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#sns TfTopicRule#sns}
	// Experimental.
	Sns interface{} `field:"optional" json:"sns" yaml:"sns"`
	// sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#sqs TfTopicRule#sqs}
	// Experimental.
	Sqs interface{} `field:"optional" json:"sqs" yaml:"sqs"`
	// step_functions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#step_functions TfTopicRule#step_functions}
	// Experimental.
	StepFunctions interface{} `field:"optional" json:"stepFunctions" yaml:"stepFunctions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#tags TfTopicRule#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#tags_all TfTopicRule#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timestream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#timestream TfTopicRule#timestream}
	// Experimental.
	Timestream interface{} `field:"optional" json:"timestream" yaml:"timestream"`
}

