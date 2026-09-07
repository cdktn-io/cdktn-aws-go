package bedrockagentcore


// Experimental.
type AwsEvaluator_LambdaConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#lambda_arn AwsEvaluator#lambda_arn}.
	// Experimental.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#lambda_timeout_in_seconds AwsEvaluator#lambda_timeout_in_seconds}.
	// Experimental.
	LambdaTimeoutInSeconds *float64 `field:"optional" json:"lambdaTimeoutInSeconds" yaml:"lambdaTimeoutInSeconds"`
}

