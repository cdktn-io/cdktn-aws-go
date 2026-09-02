package awsec2


// Experimental.
type TfAmi_EphemeralBlockDeviceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#device_name TfAmi#device_name}.
	// Experimental.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#virtual_name TfAmi#virtual_name}.
	// Experimental.
	VirtualName *string `field:"required" json:"virtualName" yaml:"virtualName"`
}

