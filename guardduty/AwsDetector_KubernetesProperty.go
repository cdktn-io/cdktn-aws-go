package guardduty


// Experimental.
type AwsDetector_KubernetesProperty struct {
	// audit_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector#audit_logs AwsDetector#audit_logs}
	// Experimental.
	AuditLogs *AwsDetector_AuditLogsProperty `field:"required" json:"auditLogs" yaml:"auditLogs"`
}

