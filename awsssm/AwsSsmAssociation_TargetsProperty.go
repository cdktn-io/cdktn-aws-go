package awsssm


// Experimental.
type AwsSsmAssociation_TargetsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_association#key AwsSsmAssociation#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_association#values AwsSsmAssociation#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

