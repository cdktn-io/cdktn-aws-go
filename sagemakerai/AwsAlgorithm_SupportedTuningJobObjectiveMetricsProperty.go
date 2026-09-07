package sagemakerai


// Experimental.
type AwsAlgorithm_SupportedTuningJobObjectiveMetricsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#metric_name AwsAlgorithm#metric_name}.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#type AwsAlgorithm#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

