package ec2


// Experimental.
type AwsSpotFleetRequest_EphemeralBlockDeviceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#device_name AwsSpotFleetRequest#device_name}.
	// Experimental.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#virtual_name AwsSpotFleetRequest#virtual_name}.
	// Experimental.
	VirtualName *string `field:"required" json:"virtualName" yaml:"virtualName"`
}

