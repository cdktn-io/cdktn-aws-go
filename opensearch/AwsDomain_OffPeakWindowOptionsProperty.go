package opensearch


// Experimental.
type AwsDomain_OffPeakWindowOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#enabled AwsDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// off_peak_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#off_peak_window AwsDomain#off_peak_window}
	// Experimental.
	OffPeakWindow *AwsDomain_OffPeakWindowProperty `field:"optional" json:"offPeakWindow" yaml:"offPeakWindow"`
}

