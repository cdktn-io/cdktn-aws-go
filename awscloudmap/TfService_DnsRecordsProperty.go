package awscloudmap


// Experimental.
type TfService_DnsRecordsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#ttl TfService#ttl}.
	// Experimental.
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#type TfService#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

