package sagemakerai


// Experimental.
type AwsWorkforce_CognitoConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#client_id AwsWorkforce#client_id}.
	// Experimental.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#user_pool AwsWorkforce#user_pool}.
	// Experimental.
	UserPool *string `field:"required" json:"userPool" yaml:"userPool"`
}

