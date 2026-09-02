package awsiotcore


// Experimental.
type TfTopicRule_ErrorActionCloudwatchMetricProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#metric_name TfTopicRule#metric_name}.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#metric_namespace TfTopicRule#metric_namespace}.
	// Experimental.
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#metric_unit TfTopicRule#metric_unit}.
	// Experimental.
	MetricUnit *string `field:"required" json:"metricUnit" yaml:"metricUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#metric_value TfTopicRule#metric_value}.
	// Experimental.
	MetricValue *string `field:"required" json:"metricValue" yaml:"metricValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#role_arn TfTopicRule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#metric_timestamp TfTopicRule#metric_timestamp}.
	// Experimental.
	MetricTimestamp *string `field:"optional" json:"metricTimestamp" yaml:"metricTimestamp"`
}

