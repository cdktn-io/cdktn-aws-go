package awsappsync


// Experimental.
type AwsAppsyncFunction_SyncConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_function#conflict_detection AwsAppsyncFunction#conflict_detection}.
	// Experimental.
	ConflictDetection *string `field:"optional" json:"conflictDetection" yaml:"conflictDetection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_function#conflict_handler AwsAppsyncFunction#conflict_handler}.
	// Experimental.
	ConflictHandler *string `field:"optional" json:"conflictHandler" yaml:"conflictHandler"`
	// lambda_conflict_handler_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_function#lambda_conflict_handler_config AwsAppsyncFunction#lambda_conflict_handler_config}
	// Experimental.
	LambdaConflictHandlerConfig *AwsAppsyncFunction_LambdaConflictHandlerConfigProperty `field:"optional" json:"lambdaConflictHandlerConfig" yaml:"lambdaConflictHandlerConfig"`
}

