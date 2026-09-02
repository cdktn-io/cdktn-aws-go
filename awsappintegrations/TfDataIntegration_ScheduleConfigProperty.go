package awsappintegrations


// Experimental.
type TfDataIntegration_ScheduleConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appintegrations_data_integration#first_execution_from TfDataIntegration#first_execution_from}.
	// Experimental.
	FirstExecutionFrom *string `field:"required" json:"firstExecutionFrom" yaml:"firstExecutionFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appintegrations_data_integration#object TfDataIntegration#object}.
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appintegrations_data_integration#schedule_expression TfDataIntegration#schedule_expression}.
	// Experimental.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
}

