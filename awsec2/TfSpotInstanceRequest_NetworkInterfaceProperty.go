package awsec2


// Experimental.
type TfSpotInstanceRequest_NetworkInterfaceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#device_index TfSpotInstanceRequest#device_index}.
	// Experimental.
	DeviceIndex *float64 `field:"required" json:"deviceIndex" yaml:"deviceIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#network_interface_id TfSpotInstanceRequest#network_interface_id}.
	// Experimental.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#delete_on_termination TfSpotInstanceRequest#delete_on_termination}.
	// Experimental.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
}

