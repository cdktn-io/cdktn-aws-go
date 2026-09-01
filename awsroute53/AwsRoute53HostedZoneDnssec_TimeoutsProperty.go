package awsroute53


// Experimental.
type AwsRoute53HostedZoneDnssec_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_hosted_zone_dnssec#create AwsRoute53HostedZoneDnssec#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_hosted_zone_dnssec#delete AwsRoute53HostedZoneDnssec#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_hosted_zone_dnssec#update AwsRoute53HostedZoneDnssec#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

