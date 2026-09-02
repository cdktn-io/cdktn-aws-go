package awsec2


// Experimental.
type TfSpotFleetRequest_EphemeralBlockDeviceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#device_name TfSpotFleetRequest#device_name}.
	// Experimental.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#virtual_name TfSpotFleetRequest#virtual_name}.
	// Experimental.
	VirtualName *string `field:"required" json:"virtualName" yaml:"virtualName"`
}

