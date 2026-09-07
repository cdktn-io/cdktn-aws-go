package route53resolver


// Experimental.
type DataAwsResolverQueryLogConfig_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_resolver_query_log_config#name DataAwsResolverQueryLogConfig#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_resolver_query_log_config#values DataAwsResolverQueryLogConfig#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

