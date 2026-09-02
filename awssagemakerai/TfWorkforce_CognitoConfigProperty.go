package awssagemakerai


// Experimental.
type TfWorkforce_CognitoConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#client_id TfWorkforce#client_id}.
	// Experimental.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#user_pool TfWorkforce#user_pool}.
	// Experimental.
	UserPool *string `field:"required" json:"userPool" yaml:"userPool"`
}

