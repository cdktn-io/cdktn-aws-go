package awsappsync


// Experimental.
type TfFunction_SyncConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_function#conflict_detection TfFunction#conflict_detection}.
	// Experimental.
	ConflictDetection *string `field:"optional" json:"conflictDetection" yaml:"conflictDetection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_function#conflict_handler TfFunction#conflict_handler}.
	// Experimental.
	ConflictHandler *string `field:"optional" json:"conflictHandler" yaml:"conflictHandler"`
	// lambda_conflict_handler_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_function#lambda_conflict_handler_config TfFunction#lambda_conflict_handler_config}
	// Experimental.
	LambdaConflictHandlerConfig *TfFunction_LambdaConflictHandlerConfigProperty `field:"optional" json:"lambdaConflictHandlerConfig" yaml:"lambdaConflictHandlerConfig"`
}

