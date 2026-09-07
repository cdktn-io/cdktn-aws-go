package ssm


// Experimental.
type AwsMaintenanceWindowTask_TaskInvocationParametersRunCommandParametersParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#name AwsMaintenanceWindowTask#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#values AwsMaintenanceWindowTask#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

