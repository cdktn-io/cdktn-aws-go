package awsroute53


// Experimental.
type AwsRoute53VpcAssociationAuthorization_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_vpc_association_authorization#create AwsRoute53VpcAssociationAuthorization#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_vpc_association_authorization#delete AwsRoute53VpcAssociationAuthorization#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_vpc_association_authorization#read AwsRoute53VpcAssociationAuthorization#read}.
	// Experimental.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

