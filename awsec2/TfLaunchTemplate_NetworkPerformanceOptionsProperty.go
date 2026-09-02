package awsec2


// Experimental.
type TfLaunchTemplate_NetworkPerformanceOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#bandwidth_weighting TfLaunchTemplate#bandwidth_weighting}.
	// Experimental.
	BandwidthWeighting *string `field:"optional" json:"bandwidthWeighting" yaml:"bandwidthWeighting"`
}

