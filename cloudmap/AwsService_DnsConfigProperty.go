package cloudmap


// Experimental.
type AwsService_DnsConfigProperty struct {
	// dns_records block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#dns_records AwsService#dns_records}
	// Experimental.
	DnsRecords interface{} `field:"required" json:"dnsRecords" yaml:"dnsRecords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#namespace_id AwsService#namespace_id}.
	// Experimental.
	NamespaceId *string `field:"required" json:"namespaceId" yaml:"namespaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#routing_policy AwsService#routing_policy}.
	// Experimental.
	RoutingPolicy *string `field:"optional" json:"routingPolicy" yaml:"routingPolicy"`
}

