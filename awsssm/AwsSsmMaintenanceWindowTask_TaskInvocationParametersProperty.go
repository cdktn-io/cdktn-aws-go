package awsssm


// Experimental.
type AwsSsmMaintenanceWindowTask_TaskInvocationParametersProperty struct {
	// automation_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#automation_parameters AwsSsmMaintenanceWindowTask#automation_parameters}
	// Experimental.
	AutomationParameters *AwsSsmMaintenanceWindowTask_AutomationParametersProperty `field:"optional" json:"automationParameters" yaml:"automationParameters"`
	// lambda_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#lambda_parameters AwsSsmMaintenanceWindowTask#lambda_parameters}
	// Experimental.
	LambdaParameters *AwsSsmMaintenanceWindowTask_LambdaParametersProperty `field:"optional" json:"lambdaParameters" yaml:"lambdaParameters"`
	// run_command_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#run_command_parameters AwsSsmMaintenanceWindowTask#run_command_parameters}
	// Experimental.
	RunCommandParameters *AwsSsmMaintenanceWindowTask_RunCommandParametersProperty `field:"optional" json:"runCommandParameters" yaml:"runCommandParameters"`
	// step_functions_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#step_functions_parameters AwsSsmMaintenanceWindowTask#step_functions_parameters}
	// Experimental.
	StepFunctionsParameters *AwsSsmMaintenanceWindowTask_StepFunctionsParametersProperty `field:"optional" json:"stepFunctionsParameters" yaml:"stepFunctionsParameters"`
}

