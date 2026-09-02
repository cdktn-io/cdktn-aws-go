package awselementalmedialive


// Experimental.
type TfChannel_TimecodeConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#source TfChannel#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#sync_threshold TfChannel#sync_threshold}.
	// Experimental.
	SyncThreshold *float64 `field:"optional" json:"syncThreshold" yaml:"syncThreshold"`
}

