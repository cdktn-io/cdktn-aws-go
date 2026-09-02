package awsec2imagebuilder


// Experimental.
type TfImage_WorkflowProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image#workflow_arn TfImage#workflow_arn}.
	// Experimental.
	WorkflowArn *string `field:"required" json:"workflowArn" yaml:"workflowArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image#on_failure TfImage#on_failure}.
	// Experimental.
	OnFailure *string `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image#parallel_group TfImage#parallel_group}.
	// Experimental.
	ParallelGroup *string `field:"optional" json:"parallelGroup" yaml:"parallelGroup"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image#parameter TfImage#parameter}
	// Experimental.
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

