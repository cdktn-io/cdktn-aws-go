package opensearch


// Experimental.
type AwsOutboundConnection_RemoteDomainInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#domain_name AwsOutboundConnection#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#owner_id AwsOutboundConnection#owner_id}.
	// Experimental.
	OwnerId *string `field:"required" json:"ownerId" yaml:"ownerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#region AwsOutboundConnection#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
}

