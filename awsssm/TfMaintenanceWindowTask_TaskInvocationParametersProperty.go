package awsssm


// Experimental.
type TfMaintenanceWindowTask_TaskInvocationParametersProperty struct {
	// automation_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#automation_parameters TfMaintenanceWindowTask#automation_parameters}
	// Experimental.
	AutomationParameters *TfMaintenanceWindowTask_AutomationParametersProperty `field:"optional" json:"automationParameters" yaml:"automationParameters"`
	// lambda_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#lambda_parameters TfMaintenanceWindowTask#lambda_parameters}
	// Experimental.
	LambdaParameters *TfMaintenanceWindowTask_LambdaParametersProperty `field:"optional" json:"lambdaParameters" yaml:"lambdaParameters"`
	// run_command_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#run_command_parameters TfMaintenanceWindowTask#run_command_parameters}
	// Experimental.
	RunCommandParameters *TfMaintenanceWindowTask_RunCommandParametersProperty `field:"optional" json:"runCommandParameters" yaml:"runCommandParameters"`
	// step_functions_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#step_functions_parameters TfMaintenanceWindowTask#step_functions_parameters}
	// Experimental.
	StepFunctionsParameters *TfMaintenanceWindowTask_StepFunctionsParametersProperty `field:"optional" json:"stepFunctionsParameters" yaml:"stepFunctionsParameters"`
}

