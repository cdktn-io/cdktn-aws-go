package awsroute53


// Experimental.
type AwsRoute53Record_GeolocationRoutingPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#continent AwsRoute53Record#continent}.
	// Experimental.
	Continent *string `field:"optional" json:"continent" yaml:"continent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#country AwsRoute53Record#country}.
	// Experimental.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#subdivision AwsRoute53Record#subdivision}.
	// Experimental.
	Subdivision *string `field:"optional" json:"subdivision" yaml:"subdivision"`
}

