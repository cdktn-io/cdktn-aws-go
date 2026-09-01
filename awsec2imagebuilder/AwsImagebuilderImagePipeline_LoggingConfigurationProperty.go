package awsec2imagebuilder


// Experimental.
type AwsImagebuilderImagePipeline_LoggingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#image_log_group_name AwsImagebuilderImagePipeline#image_log_group_name}.
	// Experimental.
	ImageLogGroupName *string `field:"optional" json:"imageLogGroupName" yaml:"imageLogGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#pipeline_log_group_name AwsImagebuilderImagePipeline#pipeline_log_group_name}.
	// Experimental.
	PipelineLogGroupName *string `field:"optional" json:"pipelineLogGroupName" yaml:"pipelineLogGroupName"`
}

