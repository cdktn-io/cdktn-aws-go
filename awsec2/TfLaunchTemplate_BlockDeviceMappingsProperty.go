package awsec2


// Experimental.
type TfLaunchTemplate_BlockDeviceMappingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#device_name TfLaunchTemplate#device_name}.
	// Experimental.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// ebs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ebs TfLaunchTemplate#ebs}
	// Experimental.
	Ebs *TfLaunchTemplate_EbsProperty `field:"optional" json:"ebs" yaml:"ebs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#no_device TfLaunchTemplate#no_device}.
	// Experimental.
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#virtual_name TfLaunchTemplate#virtual_name}.
	// Experimental.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

