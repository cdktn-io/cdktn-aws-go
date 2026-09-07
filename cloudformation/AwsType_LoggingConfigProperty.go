package cloudformation


// Experimental.
type AwsType_LoggingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_type#log_group_name AwsType#log_group_name}.
	// Experimental.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_type#log_role_arn AwsType#log_role_arn}.
	// Experimental.
	LogRoleArn *string `field:"required" json:"logRoleArn" yaml:"logRoleArn"`
}

