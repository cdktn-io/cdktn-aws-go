package awselbclassic


// Experimental.
type AwsLbSslNegotiationPolicy_AttributeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_ssl_negotiation_policy#name AwsLbSslNegotiationPolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_ssl_negotiation_policy#value AwsLbSslNegotiationPolicy#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

