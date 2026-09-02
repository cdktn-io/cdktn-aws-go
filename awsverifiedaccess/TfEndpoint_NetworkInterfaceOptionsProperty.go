package awsverifiedaccess


// Experimental.
type TfEndpoint_NetworkInterfaceOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#network_interface_id TfEndpoint#network_interface_id}.
	// Experimental.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#port TfEndpoint#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#port_range TfEndpoint#port_range}
	// Experimental.
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#protocol TfEndpoint#protocol}.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

