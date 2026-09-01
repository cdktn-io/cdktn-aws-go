package awsssm


// Experimental.
type DataAwsSsmMaintenanceWindows_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ssm_maintenance_windows#name DataAwsSsmMaintenanceWindows#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ssm_maintenance_windows#values DataAwsSsmMaintenanceWindows#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

