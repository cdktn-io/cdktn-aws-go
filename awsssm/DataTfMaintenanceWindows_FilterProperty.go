package awsssm


// Experimental.
type DataTfMaintenanceWindows_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ssm_maintenance_windows#name DataTfMaintenanceWindows#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ssm_maintenance_windows#values DataTfMaintenanceWindows#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

