package awsec2imagebuilder


// Experimental.
type TfImage_LoggingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image#log_group_name TfImage#log_group_name}.
	// Experimental.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
}

