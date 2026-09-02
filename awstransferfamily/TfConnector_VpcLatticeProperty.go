package awstransferfamily


// Experimental.
type TfConnector_VpcLatticeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_connector#resource_configuration_arn TfConnector#resource_configuration_arn}.
	// Experimental.
	ResourceConfigurationArn *string `field:"required" json:"resourceConfigurationArn" yaml:"resourceConfigurationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_connector#port_number TfConnector#port_number}.
	// Experimental.
	PortNumber *float64 `field:"optional" json:"portNumber" yaml:"portNumber"`
}

