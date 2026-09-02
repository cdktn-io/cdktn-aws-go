package awseventbridge


// Experimental.
type TfConnection_ConnectivityParametersProperty struct {
	// resource_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#resource_parameters TfConnection#resource_parameters}
	// Experimental.
	ResourceParameters *TfConnection_AuthParametersConnectivityParametersResourceParametersProperty `field:"required" json:"resourceParameters" yaml:"resourceParameters"`
}

