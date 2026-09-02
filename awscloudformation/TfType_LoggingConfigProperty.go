package awscloudformation


// Experimental.
type TfType_LoggingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_type#log_group_name TfType#log_group_name}.
	// Experimental.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_type#log_role_arn TfType#log_role_arn}.
	// Experimental.
	LogRoleArn *string `field:"required" json:"logRoleArn" yaml:"logRoleArn"`
}

