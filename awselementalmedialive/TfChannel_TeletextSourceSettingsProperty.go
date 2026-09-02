package awselementalmedialive


// Experimental.
type TfChannel_TeletextSourceSettingsProperty struct {
	// output_rectangle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#output_rectangle TfChannel#output_rectangle}
	// Experimental.
	OutputRectangle *TfChannel_OutputRectangleProperty `field:"optional" json:"outputRectangle" yaml:"outputRectangle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#page_number TfChannel#page_number}.
	// Experimental.
	PageNumber *string `field:"optional" json:"pageNumber" yaml:"pageNumber"`
}

