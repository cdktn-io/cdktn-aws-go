package awskinesisanalyticsv2


// Experimental.
type TfApplication_ApplicationCodeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#code_content_type TfApplication#code_content_type}.
	// Experimental.
	CodeContentType *string `field:"required" json:"codeContentType" yaml:"codeContentType"`
	// code_content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#code_content TfApplication#code_content}
	// Experimental.
	CodeContent *TfApplication_CodeContentProperty `field:"optional" json:"codeContent" yaml:"codeContent"`
}

