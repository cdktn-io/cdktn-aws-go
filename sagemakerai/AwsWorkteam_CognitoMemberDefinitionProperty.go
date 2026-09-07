package sagemakerai


// Experimental.
type AwsWorkteam_CognitoMemberDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workteam#client_id AwsWorkteam#client_id}.
	// Experimental.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workteam#user_group AwsWorkteam#user_group}.
	// Experimental.
	UserGroup *string `field:"required" json:"userGroup" yaml:"userGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workteam#user_pool AwsWorkteam#user_pool}.
	// Experimental.
	UserPool *string `field:"required" json:"userPool" yaml:"userPool"`
}

