package awsroute53recoveryreadiness


// Experimental.
type AwsRoute53RecoveryreadinessResourceSet_TargetResourceProperty struct {
	// nlb_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#nlb_resource AwsRoute53RecoveryreadinessResourceSet#nlb_resource}
	// Experimental.
	NlbResource *AwsRoute53RecoveryreadinessResourceSet_NlbResourceProperty `field:"optional" json:"nlbResource" yaml:"nlbResource"`
	// r53_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#r53_resource AwsRoute53RecoveryreadinessResourceSet#r53_resource}
	// Experimental.
	R53Resource *AwsRoute53RecoveryreadinessResourceSet_R53ResourceProperty `field:"optional" json:"r53Resource" yaml:"r53Resource"`
}

