package awselementalmedialive


// Experimental.
type TfChannel_AvailBlankingProperty struct {
	// avail_blanking_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#avail_blanking_image TfChannel#avail_blanking_image}
	// Experimental.
	AvailBlankingImage *TfChannel_AvailBlankingImageProperty `field:"optional" json:"availBlankingImage" yaml:"availBlankingImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#state TfChannel#state}.
	// Experimental.
	State *string `field:"optional" json:"state" yaml:"state"`
}

