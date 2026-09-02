package awscloudwatch


// Experimental.
type TfMetricStream_ExcludeFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#namespace TfMetricStream#namespace}.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#metric_names TfMetricStream#metric_names}.
	// Experimental.
	MetricNames *[]*string `field:"optional" json:"metricNames" yaml:"metricNames"`
}

