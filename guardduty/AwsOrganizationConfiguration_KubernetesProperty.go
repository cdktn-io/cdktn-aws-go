package guardduty


// Experimental.
type AwsOrganizationConfiguration_KubernetesProperty struct {
	// audit_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_organization_configuration#audit_logs AwsOrganizationConfiguration#audit_logs}
	// Experimental.
	AuditLogs *AwsOrganizationConfiguration_AuditLogsProperty `field:"required" json:"auditLogs" yaml:"auditLogs"`
}

