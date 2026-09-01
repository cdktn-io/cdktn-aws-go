package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreEvaluator_LambdaConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#lambda_arn AwsBedrockagentcoreEvaluator#lambda_arn}.
	// Experimental.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#lambda_timeout_in_seconds AwsBedrockagentcoreEvaluator#lambda_timeout_in_seconds}.
	// Experimental.
	LambdaTimeoutInSeconds *float64 `field:"optional" json:"lambdaTimeoutInSeconds" yaml:"lambdaTimeoutInSeconds"`
}

