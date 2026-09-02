package awsec2imagebuilder


// Experimental.
type TfImagePipeline_WorkflowProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#workflow_arn TfImagePipeline#workflow_arn}.
	// Experimental.
	WorkflowArn *string `field:"required" json:"workflowArn" yaml:"workflowArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#on_failure TfImagePipeline#on_failure}.
	// Experimental.
	OnFailure *string `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#parallel_group TfImagePipeline#parallel_group}.
	// Experimental.
	ParallelGroup *string `field:"optional" json:"parallelGroup" yaml:"parallelGroup"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#parameter TfImagePipeline#parameter}
	// Experimental.
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

