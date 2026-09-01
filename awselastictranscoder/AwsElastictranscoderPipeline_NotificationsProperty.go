package awselastictranscoder


// Experimental.
type AwsElastictranscoderPipeline_NotificationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#completed AwsElastictranscoderPipeline#completed}.
	// Experimental.
	Completed *string `field:"optional" json:"completed" yaml:"completed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#error AwsElastictranscoderPipeline#error}.
	// Experimental.
	Error *string `field:"optional" json:"error" yaml:"error"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#progressing AwsElastictranscoderPipeline#progressing}.
	// Experimental.
	Progressing *string `field:"optional" json:"progressing" yaml:"progressing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#warning AwsElastictranscoderPipeline#warning}.
	// Experimental.
	Warning *string `field:"optional" json:"warning" yaml:"warning"`
}

