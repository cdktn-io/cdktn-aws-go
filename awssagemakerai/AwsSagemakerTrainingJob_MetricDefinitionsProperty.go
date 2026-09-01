package awssagemakerai


// Experimental.
type AwsSagemakerTrainingJob_MetricDefinitionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#name AwsSagemakerTrainingJob#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#regex AwsSagemakerTrainingJob#regex}.
	// Experimental.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
}

