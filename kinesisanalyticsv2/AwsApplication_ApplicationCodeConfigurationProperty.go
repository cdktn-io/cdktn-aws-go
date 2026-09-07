package kinesisanalyticsv2


// Experimental.
type AwsApplication_ApplicationCodeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#code_content_type AwsApplication#code_content_type}.
	// Experimental.
	CodeContentType *string `field:"required" json:"codeContentType" yaml:"codeContentType"`
	// code_content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#code_content AwsApplication#code_content}
	// Experimental.
	CodeContent *AwsApplication_CodeContentProperty `field:"optional" json:"codeContent" yaml:"codeContent"`
}

