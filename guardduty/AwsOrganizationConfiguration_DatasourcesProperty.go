package guardduty


// Experimental.
type AwsOrganizationConfiguration_DatasourcesProperty struct {
	// kubernetes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_organization_configuration#kubernetes AwsOrganizationConfiguration#kubernetes}
	// Experimental.
	Kubernetes *AwsOrganizationConfiguration_KubernetesProperty `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// malware_protection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_organization_configuration#malware_protection AwsOrganizationConfiguration#malware_protection}
	// Experimental.
	MalwareProtection *AwsOrganizationConfiguration_MalwareProtectionProperty `field:"optional" json:"malwareProtection" yaml:"malwareProtection"`
	// s3_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_organization_configuration#s3_logs AwsOrganizationConfiguration#s3_logs}
	// Experimental.
	S3Logs *AwsOrganizationConfiguration_S3LogsProperty `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

