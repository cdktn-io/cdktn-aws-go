package awssagemakerai


// Experimental.
type AwsSagemakerWorkteam_CognitoMemberDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workteam#client_id AwsSagemakerWorkteam#client_id}.
	// Experimental.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workteam#user_group AwsSagemakerWorkteam#user_group}.
	// Experimental.
	UserGroup *string `field:"required" json:"userGroup" yaml:"userGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workteam#user_pool AwsSagemakerWorkteam#user_pool}.
	// Experimental.
	UserPool *string `field:"required" json:"userPool" yaml:"userPool"`
}

