package awsec2


// Experimental.
type TfLaunchTemplate_LicenseSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#license_configuration_arn TfLaunchTemplate#license_configuration_arn}.
	// Experimental.
	LicenseConfigurationArn *string `field:"required" json:"licenseConfigurationArn" yaml:"licenseConfigurationArn"`
}

