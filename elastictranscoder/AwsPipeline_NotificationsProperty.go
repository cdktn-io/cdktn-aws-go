package elastictranscoder


// Experimental.
type AwsPipeline_NotificationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#completed AwsPipeline#completed}.
	// Experimental.
	Completed *string `field:"optional" json:"completed" yaml:"completed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#error AwsPipeline#error}.
	// Experimental.
	Error *string `field:"optional" json:"error" yaml:"error"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#progressing AwsPipeline#progressing}.
	// Experimental.
	Progressing *string `field:"optional" json:"progressing" yaml:"progressing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#warning AwsPipeline#warning}.
	// Experimental.
	Warning *string `field:"optional" json:"warning" yaml:"warning"`
}

