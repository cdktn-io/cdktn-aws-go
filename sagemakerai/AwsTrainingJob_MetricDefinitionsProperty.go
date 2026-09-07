package sagemakerai


// Experimental.
type AwsTrainingJob_MetricDefinitionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#name AwsTrainingJob#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#regex AwsTrainingJob#regex}.
	// Experimental.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
}

