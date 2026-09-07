package licensemanager


// Experimental.
type DataAwsReceivedLicenses_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/licensemanager_received_licenses#name DataAwsReceivedLicenses#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/licensemanager_received_licenses#values DataAwsReceivedLicenses#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

