package awsguardduty


// Experimental.
type TfDetector_ScanEc2InstanceWithFindingsProperty struct {
	// ebs_volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector#ebs_volumes TfDetector#ebs_volumes}
	// Experimental.
	EbsVolumes *TfDetector_EbsVolumesProperty `field:"required" json:"ebsVolumes" yaml:"ebsVolumes"`
}

