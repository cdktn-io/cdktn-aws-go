package route53


// Experimental.
type AwsRecord_WeightedRoutingPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#weight AwsRecord#weight}.
	// Experimental.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

