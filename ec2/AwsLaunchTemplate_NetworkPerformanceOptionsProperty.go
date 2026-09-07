package ec2


// Experimental.
type AwsLaunchTemplate_NetworkPerformanceOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#bandwidth_weighting AwsLaunchTemplate#bandwidth_weighting}.
	// Experimental.
	BandwidthWeighting *string `field:"optional" json:"bandwidthWeighting" yaml:"bandwidthWeighting"`
}

