package awselementalmedialive


// Experimental.
type TfChannel_InputChannelLevelsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#gain TfChannel#gain}.
	// Experimental.
	Gain *float64 `field:"required" json:"gain" yaml:"gain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_channel TfChannel#input_channel}.
	// Experimental.
	InputChannel *float64 `field:"required" json:"inputChannel" yaml:"inputChannel"`
}

