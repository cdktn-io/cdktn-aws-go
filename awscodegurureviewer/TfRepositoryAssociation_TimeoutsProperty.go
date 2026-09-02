package awscodegurureviewer


// Experimental.
type TfRepositoryAssociation_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#create TfRepositoryAssociation#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#delete TfRepositoryAssociation#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#update TfRepositoryAssociation#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

