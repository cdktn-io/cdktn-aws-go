package awsssm


// Experimental.
type TfMaintenanceWindowTask_TargetsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#key TfMaintenanceWindowTask#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#values TfMaintenanceWindowTask#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

