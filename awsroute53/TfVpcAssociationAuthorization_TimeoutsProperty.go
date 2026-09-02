package awsroute53


// Experimental.
type TfVpcAssociationAuthorization_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_vpc_association_authorization#create TfVpcAssociationAuthorization#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_vpc_association_authorization#delete TfVpcAssociationAuthorization#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_vpc_association_authorization#read TfVpcAssociationAuthorization#read}.
	// Experimental.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

