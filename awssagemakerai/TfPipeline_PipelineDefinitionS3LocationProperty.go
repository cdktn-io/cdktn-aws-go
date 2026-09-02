package awssagemakerai


// Experimental.
type TfPipeline_PipelineDefinitionS3LocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_pipeline#bucket TfPipeline#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_pipeline#object_key TfPipeline#object_key}.
	// Experimental.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_pipeline#version_id TfPipeline#version_id}.
	// Experimental.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

