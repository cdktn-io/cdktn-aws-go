package awsecr


// Experimental.
type TfRegistryScanningConfiguration_RuleProperty struct {
	// repository_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_registry_scanning_configuration#repository_filter TfRegistryScanningConfiguration#repository_filter}
	// Experimental.
	RepositoryFilter interface{} `field:"required" json:"repositoryFilter" yaml:"repositoryFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_registry_scanning_configuration#scan_frequency TfRegistryScanningConfiguration#scan_frequency}.
	// Experimental.
	ScanFrequency *string `field:"required" json:"scanFrequency" yaml:"scanFrequency"`
}

