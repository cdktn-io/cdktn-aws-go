package awsecr


// Experimental.
type AwsEcrRegistryScanningConfiguration_RuleProperty struct {
	// repository_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_registry_scanning_configuration#repository_filter AwsEcrRegistryScanningConfiguration#repository_filter}
	// Experimental.
	RepositoryFilter interface{} `field:"required" json:"repositoryFilter" yaml:"repositoryFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_registry_scanning_configuration#scan_frequency AwsEcrRegistryScanningConfiguration#scan_frequency}.
	// Experimental.
	ScanFrequency *string `field:"required" json:"scanFrequency" yaml:"scanFrequency"`
}

