package awsroute53


// Experimental.
type TfRecord_GeolocationRoutingPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#continent TfRecord#continent}.
	// Experimental.
	Continent *string `field:"optional" json:"continent" yaml:"continent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#country TfRecord#country}.
	// Experimental.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#subdivision TfRecord#subdivision}.
	// Experimental.
	Subdivision *string `field:"optional" json:"subdivision" yaml:"subdivision"`
}

