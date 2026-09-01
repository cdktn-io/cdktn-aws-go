package awsguardduty


// Experimental.
type AwsGuarddutyOrganizationConfiguration_ScanEc2InstanceWithFindingsProperty struct {
	// ebs_volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_organization_configuration#ebs_volumes AwsGuarddutyOrganizationConfiguration#ebs_volumes}
	// Experimental.
	EbsVolumes *AwsGuarddutyOrganizationConfiguration_EbsVolumesProperty `field:"required" json:"ebsVolumes" yaml:"ebsVolumes"`
}

