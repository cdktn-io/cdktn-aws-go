package awsguardduty


// Experimental.
type TfOrganizationConfiguration_KubernetesProperty struct {
	// audit_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_organization_configuration#audit_logs TfOrganizationConfiguration#audit_logs}
	// Experimental.
	AuditLogs *TfOrganizationConfiguration_AuditLogsProperty `field:"required" json:"auditLogs" yaml:"auditLogs"`
}

