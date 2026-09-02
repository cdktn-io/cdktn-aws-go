package awsec2


// Experimental.
type TfInstance_PrimaryNetworkInterfaceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#network_interface_id TfInstance#network_interface_id}.
	// Experimental.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
}

