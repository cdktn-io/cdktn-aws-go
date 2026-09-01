package awscloudwatch


// Experimental.
type AwsCloudwatchMetricStream_StatisticsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#additional_statistics AwsCloudwatchMetricStream#additional_statistics}.
	// Experimental.
	AdditionalStatistics *[]*string `field:"required" json:"additionalStatistics" yaml:"additionalStatistics"`
	// include_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#include_metric AwsCloudwatchMetricStream#include_metric}
	// Experimental.
	IncludeMetric interface{} `field:"required" json:"includeMetric" yaml:"includeMetric"`
}

