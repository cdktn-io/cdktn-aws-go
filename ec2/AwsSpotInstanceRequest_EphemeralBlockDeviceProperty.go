package ec2


// Experimental.
type AwsSpotInstanceRequest_EphemeralBlockDeviceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#device_name AwsSpotInstanceRequest#device_name}.
	// Experimental.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#no_device AwsSpotInstanceRequest#no_device}.
	// Experimental.
	NoDevice interface{} `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#virtual_name AwsSpotInstanceRequest#virtual_name}.
	// Experimental.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

