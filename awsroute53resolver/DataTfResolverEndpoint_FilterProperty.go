package awsroute53resolver


// Experimental.
type DataTfResolverEndpoint_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_resolver_endpoint#name DataTfResolverEndpoint#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_resolver_endpoint#values DataTfResolverEndpoint#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

