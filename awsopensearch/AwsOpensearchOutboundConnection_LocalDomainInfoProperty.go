package awsopensearch


// Experimental.
type AwsOpensearchOutboundConnection_LocalDomainInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#domain_name AwsOpensearchOutboundConnection#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#owner_id AwsOpensearchOutboundConnection#owner_id}.
	// Experimental.
	OwnerId *string `field:"required" json:"ownerId" yaml:"ownerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#region AwsOpensearchOutboundConnection#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
}

