package route53


// Experimental.
type AwsHostedZoneDnssec_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_hosted_zone_dnssec#create AwsHostedZoneDnssec#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_hosted_zone_dnssec#delete AwsHostedZoneDnssec#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_hosted_zone_dnssec#update AwsHostedZoneDnssec#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

