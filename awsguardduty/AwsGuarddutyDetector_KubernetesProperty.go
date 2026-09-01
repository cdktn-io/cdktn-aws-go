package awsguardduty


// Experimental.
type AwsGuarddutyDetector_KubernetesProperty struct {
	// audit_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector#audit_logs AwsGuarddutyDetector#audit_logs}
	// Experimental.
	AuditLogs *AwsGuarddutyDetector_AuditLogsProperty `field:"required" json:"auditLogs" yaml:"auditLogs"`
}

