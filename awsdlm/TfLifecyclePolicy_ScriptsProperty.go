package awsdlm


// Experimental.
type TfLifecyclePolicy_ScriptsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#execution_handler TfLifecyclePolicy#execution_handler}.
	// Experimental.
	ExecutionHandler *string `field:"required" json:"executionHandler" yaml:"executionHandler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#execute_operation_on_script_failure TfLifecyclePolicy#execute_operation_on_script_failure}.
	// Experimental.
	ExecuteOperationOnScriptFailure interface{} `field:"optional" json:"executeOperationOnScriptFailure" yaml:"executeOperationOnScriptFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#execution_handler_service TfLifecyclePolicy#execution_handler_service}.
	// Experimental.
	ExecutionHandlerService *string `field:"optional" json:"executionHandlerService" yaml:"executionHandlerService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#execution_timeout TfLifecyclePolicy#execution_timeout}.
	// Experimental.
	ExecutionTimeout *float64 `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#maximum_retry_count TfLifecyclePolicy#maximum_retry_count}.
	// Experimental.
	MaximumRetryCount *float64 `field:"optional" json:"maximumRetryCount" yaml:"maximumRetryCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#stages TfLifecyclePolicy#stages}.
	// Experimental.
	Stages *[]*string `field:"optional" json:"stages" yaml:"stages"`
}

