package awsguardduty


// Experimental.
type AwsGuarddutyDetector_DatasourcesProperty struct {
	// kubernetes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector#kubernetes AwsGuarddutyDetector#kubernetes}
	// Experimental.
	Kubernetes *AwsGuarddutyDetector_KubernetesProperty `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// malware_protection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector#malware_protection AwsGuarddutyDetector#malware_protection}
	// Experimental.
	MalwareProtection *AwsGuarddutyDetector_MalwareProtectionProperty `field:"optional" json:"malwareProtection" yaml:"malwareProtection"`
	// s3_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector#s3_logs AwsGuarddutyDetector#s3_logs}
	// Experimental.
	S3Logs *AwsGuarddutyDetector_S3LogsProperty `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

