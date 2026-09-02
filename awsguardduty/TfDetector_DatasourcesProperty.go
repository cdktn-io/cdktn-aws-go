package awsguardduty


// Experimental.
type TfDetector_DatasourcesProperty struct {
	// kubernetes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector#kubernetes TfDetector#kubernetes}
	// Experimental.
	Kubernetes *TfDetector_KubernetesProperty `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// malware_protection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector#malware_protection TfDetector#malware_protection}
	// Experimental.
	MalwareProtection *TfDetector_MalwareProtectionProperty `field:"optional" json:"malwareProtection" yaml:"malwareProtection"`
	// s3_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector#s3_logs TfDetector#s3_logs}
	// Experimental.
	S3Logs *TfDetector_S3LogsProperty `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

