package awslicensemanager


// Experimental.
type DataAwsLicensemanagerGrants_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/licensemanager_grants#name DataAwsLicensemanagerGrants#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/licensemanager_grants#values DataAwsLicensemanagerGrants#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

