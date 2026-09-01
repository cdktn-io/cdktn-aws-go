package awscloudmap


// Experimental.
type AwsServiceDiscoveryService_DnsRecordsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#ttl AwsServiceDiscoveryService#ttl}.
	// Experimental.
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#type AwsServiceDiscoveryService#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

