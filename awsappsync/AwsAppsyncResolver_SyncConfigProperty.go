package awsappsync


// Experimental.
type AwsAppsyncResolver_SyncConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#conflict_detection AwsAppsyncResolver#conflict_detection}.
	// Experimental.
	ConflictDetection *string `field:"optional" json:"conflictDetection" yaml:"conflictDetection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#conflict_handler AwsAppsyncResolver#conflict_handler}.
	// Experimental.
	ConflictHandler *string `field:"optional" json:"conflictHandler" yaml:"conflictHandler"`
	// lambda_conflict_handler_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#lambda_conflict_handler_config AwsAppsyncResolver#lambda_conflict_handler_config}
	// Experimental.
	LambdaConflictHandlerConfig *AwsAppsyncResolver_LambdaConflictHandlerConfigProperty `field:"optional" json:"lambdaConflictHandlerConfig" yaml:"lambdaConflictHandlerConfig"`
}

