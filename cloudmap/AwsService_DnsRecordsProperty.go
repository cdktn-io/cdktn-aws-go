package cloudmap


// Experimental.
type AwsService_DnsRecordsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#ttl AwsService#ttl}.
	// Experimental.
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#type AwsService#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

