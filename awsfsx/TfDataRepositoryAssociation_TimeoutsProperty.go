package awsfsx


// Experimental.
type TfDataRepositoryAssociation_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_data_repository_association#create TfDataRepositoryAssociation#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_data_repository_association#delete TfDataRepositoryAssociation#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_data_repository_association#update TfDataRepositoryAssociation#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

