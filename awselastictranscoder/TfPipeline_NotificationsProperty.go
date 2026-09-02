package awselastictranscoder


// Experimental.
type TfPipeline_NotificationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#completed TfPipeline#completed}.
	// Experimental.
	Completed *string `field:"optional" json:"completed" yaml:"completed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#error TfPipeline#error}.
	// Experimental.
	Error *string `field:"optional" json:"error" yaml:"error"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#progressing TfPipeline#progressing}.
	// Experimental.
	Progressing *string `field:"optional" json:"progressing" yaml:"progressing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#warning TfPipeline#warning}.
	// Experimental.
	Warning *string `field:"optional" json:"warning" yaml:"warning"`
}

