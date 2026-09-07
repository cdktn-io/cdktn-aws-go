package appsync


// Experimental.
type AwsFunction_SyncConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_function#conflict_detection AwsFunction#conflict_detection}.
	// Experimental.
	ConflictDetection *string `field:"optional" json:"conflictDetection" yaml:"conflictDetection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_function#conflict_handler AwsFunction#conflict_handler}.
	// Experimental.
	ConflictHandler *string `field:"optional" json:"conflictHandler" yaml:"conflictHandler"`
	// lambda_conflict_handler_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_function#lambda_conflict_handler_config AwsFunction#lambda_conflict_handler_config}
	// Experimental.
	LambdaConflictHandlerConfig *AwsFunction_LambdaConflictHandlerConfigProperty `field:"optional" json:"lambdaConflictHandlerConfig" yaml:"lambdaConflictHandlerConfig"`
}

