package route53


// Experimental.
type AwsRecord_GeolocationRoutingPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#continent AwsRecord#continent}.
	// Experimental.
	Continent *string `field:"optional" json:"continent" yaml:"continent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#country AwsRecord#country}.
	// Experimental.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#subdivision AwsRecord#subdivision}.
	// Experimental.
	Subdivision *string `field:"optional" json:"subdivision" yaml:"subdivision"`
}

