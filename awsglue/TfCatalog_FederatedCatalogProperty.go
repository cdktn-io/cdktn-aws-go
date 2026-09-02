package awsglue


// Experimental.
type TfCatalog_FederatedCatalogProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#connection_name TfCatalog#connection_name}.
	// Experimental.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#connection_type TfCatalog#connection_type}.
	// Experimental.
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#identifier TfCatalog#identifier}.
	// Experimental.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

