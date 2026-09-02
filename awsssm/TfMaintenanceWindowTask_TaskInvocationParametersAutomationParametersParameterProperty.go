package awsssm


// Experimental.
type TfMaintenanceWindowTask_TaskInvocationParametersAutomationParametersParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#name TfMaintenanceWindowTask#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#values TfMaintenanceWindowTask#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

