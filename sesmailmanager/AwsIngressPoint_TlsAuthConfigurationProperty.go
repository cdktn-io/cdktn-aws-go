package sesmailmanager


// Experimental.
type AwsIngressPoint_TlsAuthConfigurationProperty struct {
	// trust_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#trust_store AwsIngressPoint#trust_store}
	// Experimental.
	TrustStore interface{} `field:"optional" json:"trustStore" yaml:"trustStore"`
}

