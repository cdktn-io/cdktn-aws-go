package awsguardduty


// Experimental.
type AwsGuarddutyDetector_ScanEc2InstanceWithFindingsProperty struct {
	// ebs_volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector#ebs_volumes AwsGuarddutyDetector#ebs_volumes}
	// Experimental.
	EbsVolumes *AwsGuarddutyDetector_EbsVolumesProperty `field:"required" json:"ebsVolumes" yaml:"ebsVolumes"`
}

