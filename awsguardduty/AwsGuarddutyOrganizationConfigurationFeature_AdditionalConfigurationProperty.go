package awsguardduty


// Experimental.
type AwsGuarddutyOrganizationConfigurationFeature_AdditionalConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_organization_configuration_feature#auto_enable AwsGuarddutyOrganizationConfigurationFeature#auto_enable}.
	// Experimental.
	AutoEnable *string `field:"required" json:"autoEnable" yaml:"autoEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_organization_configuration_feature#name AwsGuarddutyOrganizationConfigurationFeature#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

