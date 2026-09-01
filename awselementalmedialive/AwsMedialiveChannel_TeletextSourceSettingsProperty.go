package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_TeletextSourceSettingsProperty struct {
	// output_rectangle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#output_rectangle AwsMedialiveChannel#output_rectangle}
	// Experimental.
	OutputRectangle *AwsMedialiveChannel_OutputRectangleProperty `field:"optional" json:"outputRectangle" yaml:"outputRectangle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#page_number AwsMedialiveChannel#page_number}.
	// Experimental.
	PageNumber *string `field:"optional" json:"pageNumber" yaml:"pageNumber"`
}

