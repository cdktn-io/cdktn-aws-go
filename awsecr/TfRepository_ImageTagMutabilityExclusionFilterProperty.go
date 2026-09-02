package awsecr


// Experimental.
type TfRepository_ImageTagMutabilityExclusionFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository#filter TfRepository#filter}.
	// Experimental.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository#filter_type TfRepository#filter_type}.
	// Experimental.
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
}

