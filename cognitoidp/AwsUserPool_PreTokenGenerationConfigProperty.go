package cognitoidp


// Experimental.
type AwsUserPool_PreTokenGenerationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#lambda_arn AwsUserPool#lambda_arn}.
	// Experimental.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#lambda_version AwsUserPool#lambda_version}.
	// Experimental.
	LambdaVersion *string `field:"required" json:"lambdaVersion" yaml:"lambdaVersion"`
}

