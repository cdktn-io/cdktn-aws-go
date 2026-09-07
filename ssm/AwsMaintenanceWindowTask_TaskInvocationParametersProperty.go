package ssm


// Experimental.
type AwsMaintenanceWindowTask_TaskInvocationParametersProperty struct {
	// automation_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#automation_parameters AwsMaintenanceWindowTask#automation_parameters}
	// Experimental.
	AutomationParameters *AwsMaintenanceWindowTask_AutomationParametersProperty `field:"optional" json:"automationParameters" yaml:"automationParameters"`
	// lambda_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#lambda_parameters AwsMaintenanceWindowTask#lambda_parameters}
	// Experimental.
	LambdaParameters *AwsMaintenanceWindowTask_LambdaParametersProperty `field:"optional" json:"lambdaParameters" yaml:"lambdaParameters"`
	// run_command_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#run_command_parameters AwsMaintenanceWindowTask#run_command_parameters}
	// Experimental.
	RunCommandParameters *AwsMaintenanceWindowTask_RunCommandParametersProperty `field:"optional" json:"runCommandParameters" yaml:"runCommandParameters"`
	// step_functions_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#step_functions_parameters AwsMaintenanceWindowTask#step_functions_parameters}
	// Experimental.
	StepFunctionsParameters *AwsMaintenanceWindowTask_StepFunctionsParametersProperty `field:"optional" json:"stepFunctionsParameters" yaml:"stepFunctionsParameters"`
}

