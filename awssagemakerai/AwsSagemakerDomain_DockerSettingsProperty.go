package awssagemakerai


// Experimental.
type AwsSagemakerDomain_DockerSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#enable_docker_access AwsSagemakerDomain#enable_docker_access}.
	// Experimental.
	EnableDockerAccess *string `field:"optional" json:"enableDockerAccess" yaml:"enableDockerAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#vpc_only_trusted_accounts AwsSagemakerDomain#vpc_only_trusted_accounts}.
	// Experimental.
	VpcOnlyTrustedAccounts *[]*string `field:"optional" json:"vpcOnlyTrustedAccounts" yaml:"vpcOnlyTrustedAccounts"`
}

