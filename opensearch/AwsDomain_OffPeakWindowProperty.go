package opensearch


// Experimental.
type AwsDomain_OffPeakWindowProperty struct {
	// window_start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#window_start_time AwsDomain#window_start_time}
	// Experimental.
	WindowStartTime *AwsDomain_WindowStartTimeProperty `field:"optional" json:"windowStartTime" yaml:"windowStartTime"`
}

