package cloudwatch


// Experimental.
type AwsMetricStream_IncludeFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#namespace AwsMetricStream#namespace}.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#metric_names AwsMetricStream#metric_names}.
	// Experimental.
	MetricNames *[]*string `field:"optional" json:"metricNames" yaml:"metricNames"`
}

