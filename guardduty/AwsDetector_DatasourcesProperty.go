package guardduty


// Experimental.
type AwsDetector_DatasourcesProperty struct {
	// kubernetes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector#kubernetes AwsDetector#kubernetes}
	// Experimental.
	Kubernetes *AwsDetector_KubernetesProperty `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// malware_protection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector#malware_protection AwsDetector#malware_protection}
	// Experimental.
	MalwareProtection *AwsDetector_MalwareProtectionProperty `field:"optional" json:"malwareProtection" yaml:"malwareProtection"`
	// s3_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector#s3_logs AwsDetector#s3_logs}
	// Experimental.
	S3Logs *AwsDetector_S3LogsProperty `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

