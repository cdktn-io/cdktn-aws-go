package awsguardduty


// Experimental.
type AwsGuarddutyOrganizationConfiguration_DatasourcesProperty struct {
	// kubernetes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_organization_configuration#kubernetes AwsGuarddutyOrganizationConfiguration#kubernetes}
	// Experimental.
	Kubernetes *AwsGuarddutyOrganizationConfiguration_KubernetesProperty `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// malware_protection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_organization_configuration#malware_protection AwsGuarddutyOrganizationConfiguration#malware_protection}
	// Experimental.
	MalwareProtection *AwsGuarddutyOrganizationConfiguration_MalwareProtectionProperty `field:"optional" json:"malwareProtection" yaml:"malwareProtection"`
	// s3_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_organization_configuration#s3_logs AwsGuarddutyOrganizationConfiguration#s3_logs}
	// Experimental.
	S3Logs *AwsGuarddutyOrganizationConfiguration_S3LogsProperty `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

