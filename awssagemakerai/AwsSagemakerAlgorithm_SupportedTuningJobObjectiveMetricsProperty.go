package awssagemakerai


// Experimental.
type AwsSagemakerAlgorithm_SupportedTuningJobObjectiveMetricsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#metric_name AwsSagemakerAlgorithm#metric_name}.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#type AwsSagemakerAlgorithm#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

