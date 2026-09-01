package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_AvailBlankingProperty struct {
	// avail_blanking_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#avail_blanking_image AwsMedialiveChannel#avail_blanking_image}
	// Experimental.
	AvailBlankingImage *AwsMedialiveChannel_AvailBlankingImageProperty `field:"optional" json:"availBlankingImage" yaml:"availBlankingImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#state AwsMedialiveChannel#state}.
	// Experimental.
	State *string `field:"optional" json:"state" yaml:"state"`
}

