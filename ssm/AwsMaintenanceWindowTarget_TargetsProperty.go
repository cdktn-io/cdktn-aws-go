package ssm


// Experimental.
type AwsMaintenanceWindowTarget_TargetsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_target#key AwsMaintenanceWindowTarget#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_target#values AwsMaintenanceWindowTarget#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

