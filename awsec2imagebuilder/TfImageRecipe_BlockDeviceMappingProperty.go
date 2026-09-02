package awsec2imagebuilder


// Experimental.
type TfImageRecipe_BlockDeviceMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_recipe#device_name TfImageRecipe#device_name}.
	// Experimental.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// ebs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_recipe#ebs TfImageRecipe#ebs}
	// Experimental.
	Ebs *TfImageRecipe_EbsProperty `field:"optional" json:"ebs" yaml:"ebs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_recipe#no_device TfImageRecipe#no_device}.
	// Experimental.
	NoDevice interface{} `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_recipe#virtual_name TfImageRecipe#virtual_name}.
	// Experimental.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

