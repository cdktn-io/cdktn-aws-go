package awscognitoidp


// Experimental.
type TfUserPool_CustomEmailSenderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#lambda_arn TfUserPool#lambda_arn}.
	// Experimental.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#lambda_version TfUserPool#lambda_version}.
	// Experimental.
	LambdaVersion *string `field:"required" json:"lambdaVersion" yaml:"lambdaVersion"`
}

