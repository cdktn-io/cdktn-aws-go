package awsappsync


// Experimental.
type TfResolver_SyncConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#conflict_detection TfResolver#conflict_detection}.
	// Experimental.
	ConflictDetection *string `field:"optional" json:"conflictDetection" yaml:"conflictDetection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#conflict_handler TfResolver#conflict_handler}.
	// Experimental.
	ConflictHandler *string `field:"optional" json:"conflictHandler" yaml:"conflictHandler"`
	// lambda_conflict_handler_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#lambda_conflict_handler_config TfResolver#lambda_conflict_handler_config}
	// Experimental.
	LambdaConflictHandlerConfig *TfResolver_LambdaConflictHandlerConfigProperty `field:"optional" json:"lambdaConflictHandlerConfig" yaml:"lambdaConflictHandlerConfig"`
}

