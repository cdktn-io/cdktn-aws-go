package awseks


// Experimental.
type TfAddon_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#create TfAddon#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#delete TfAddon#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#update TfAddon#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

