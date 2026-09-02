package awselementalmedialive


// Experimental.
type TfChannel_AudioTrackSelectionProperty struct {
	// tracks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#tracks TfChannel#tracks}
	// Experimental.
	Tracks interface{} `field:"required" json:"tracks" yaml:"tracks"`
	// dolby_e_decode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#dolby_e_decode TfChannel#dolby_e_decode}
	// Experimental.
	DolbyEDecode *TfChannel_DolbyEDecodeProperty `field:"optional" json:"dolbyEDecode" yaml:"dolbyEDecode"`
}

