package guardduty


// Experimental.
type AwsOrganizationConfiguration_ScanEc2InstanceWithFindingsProperty struct {
	// ebs_volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_organization_configuration#ebs_volumes AwsOrganizationConfiguration#ebs_volumes}
	// Experimental.
	EbsVolumes *AwsOrganizationConfiguration_EbsVolumesProperty `field:"required" json:"ebsVolumes" yaml:"ebsVolumes"`
}

