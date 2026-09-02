package awsssm


// Experimental.
type TfAssociation_TargetsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_association#key TfAssociation#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_association#values TfAssociation#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

