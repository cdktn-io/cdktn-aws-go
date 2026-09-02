package awscloudwatch


// Experimental.
type TfMetricStream_IncludeMetricProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#metric_name TfMetricStream#metric_name}.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#namespace TfMetricStream#namespace}.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

