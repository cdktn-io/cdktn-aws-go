package awsecr


// Experimental.
type TfRegistryScanningConfiguration_RepositoryFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_registry_scanning_configuration#filter TfRegistryScanningConfiguration#filter}.
	// Experimental.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_registry_scanning_configuration#filter_type TfRegistryScanningConfiguration#filter_type}.
	// Experimental.
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
}

