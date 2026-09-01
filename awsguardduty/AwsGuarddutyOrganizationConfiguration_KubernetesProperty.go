package awsguardduty


// Experimental.
type AwsGuarddutyOrganizationConfiguration_KubernetesProperty struct {
	// audit_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_organization_configuration#audit_logs AwsGuarddutyOrganizationConfiguration#audit_logs}
	// Experimental.
	AuditLogs *AwsGuarddutyOrganizationConfiguration_AuditLogsProperty `field:"required" json:"auditLogs" yaml:"auditLogs"`
}

