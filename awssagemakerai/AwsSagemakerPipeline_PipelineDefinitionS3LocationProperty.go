package awssagemakerai


// Experimental.
type AwsSagemakerPipeline_PipelineDefinitionS3LocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_pipeline#bucket AwsSagemakerPipeline#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_pipeline#object_key AwsSagemakerPipeline#object_key}.
	// Experimental.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_pipeline#version_id AwsSagemakerPipeline#version_id}.
	// Experimental.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

