package awsroute53


// Experimental.
type TfZoneAssociation_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_zone_association#create TfZoneAssociation#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_zone_association#delete TfZoneAssociation#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

