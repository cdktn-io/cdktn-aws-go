package route53recoveryreadiness


// Experimental.
type AwsResourceSet_TargetResourceProperty struct {
	// nlb_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#nlb_resource AwsResourceSet#nlb_resource}
	// Experimental.
	NlbResource *AwsResourceSet_NlbResourceProperty `field:"optional" json:"nlbResource" yaml:"nlbResource"`
	// r53_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#r53_resource AwsResourceSet#r53_resource}
	// Experimental.
	R53Resource *AwsResourceSet_R53ResourceProperty `field:"optional" json:"r53Resource" yaml:"r53Resource"`
}

