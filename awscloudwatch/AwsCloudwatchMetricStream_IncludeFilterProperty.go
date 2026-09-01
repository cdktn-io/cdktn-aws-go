package awscloudwatch


// Experimental.
type AwsCloudwatchMetricStream_IncludeFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#namespace AwsCloudwatchMetricStream#namespace}.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#metric_names AwsCloudwatchMetricStream#metric_names}.
	// Experimental.
	MetricNames *[]*string `field:"optional" json:"metricNames" yaml:"metricNames"`
}

