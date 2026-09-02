package awsecr


// Experimental.
type TfRepositoryCreationTemplate_ImageTagMutabilityExclusionFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository_creation_template#filter TfRepositoryCreationTemplate#filter}.
	// Experimental.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository_creation_template#filter_type TfRepositoryCreationTemplate#filter_type}.
	// Experimental.
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
}

