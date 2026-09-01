package awscloudformation


// Experimental.
type AwsCloudformationType_LoggingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_type#log_group_name AwsCloudformationType#log_group_name}.
	// Experimental.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_type#log_role_arn AwsCloudformationType#log_role_arn}.
	// Experimental.
	LogRoleArn *string `field:"required" json:"logRoleArn" yaml:"logRoleArn"`
}

