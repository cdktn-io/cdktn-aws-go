package ssm


// Experimental.
type AwsMaintenanceWindowTask_TargetsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#key AwsMaintenanceWindowTask#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#values AwsMaintenanceWindowTask#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

