package awsguardduty


// Experimental.
type TfDetector_KubernetesProperty struct {
	// audit_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector#audit_logs TfDetector#audit_logs}
	// Experimental.
	AuditLogs *TfDetector_AuditLogsProperty `field:"required" json:"auditLogs" yaml:"auditLogs"`
}

