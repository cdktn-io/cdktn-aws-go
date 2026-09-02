package awsappstream20


// Experimental.
type TfStack_StorageConnectorsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#connector_type TfStack#connector_type}.
	// Experimental.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#domains TfStack#domains}.
	// Experimental.
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#resource_identifier TfStack#resource_identifier}.
	// Experimental.
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
}

