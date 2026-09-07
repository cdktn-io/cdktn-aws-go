package elementalmedialive


// Experimental.
type AwsChannel_HlsInputSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#bandwidth AwsChannel#bandwidth}.
	// Experimental.
	Bandwidth *float64 `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#buffer_segments AwsChannel#buffer_segments}.
	// Experimental.
	BufferSegments *float64 `field:"optional" json:"bufferSegments" yaml:"bufferSegments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#retries AwsChannel#retries}.
	// Experimental.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#retry_interval AwsChannel#retry_interval}.
	// Experimental.
	RetryInterval *float64 `field:"optional" json:"retryInterval" yaml:"retryInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte35_source AwsChannel#scte35_source}.
	// Experimental.
	Scte35Source *string `field:"optional" json:"scte35Source" yaml:"scte35Source"`
}

