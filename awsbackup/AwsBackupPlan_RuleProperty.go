package awsbackup


// Experimental.
type AwsBackupPlan_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#rule_name AwsBackupPlan#rule_name}.
	// Experimental.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#target_vault_name AwsBackupPlan#target_vault_name}.
	// Experimental.
	TargetVaultName *string `field:"required" json:"targetVaultName" yaml:"targetVaultName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#completion_window AwsBackupPlan#completion_window}.
	// Experimental.
	CompletionWindow *float64 `field:"optional" json:"completionWindow" yaml:"completionWindow"`
	// copy_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#copy_action AwsBackupPlan#copy_action}
	// Experimental.
	CopyAction interface{} `field:"optional" json:"copyAction" yaml:"copyAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#enable_continuous_backup AwsBackupPlan#enable_continuous_backup}.
	// Experimental.
	EnableContinuousBackup interface{} `field:"optional" json:"enableContinuousBackup" yaml:"enableContinuousBackup"`
	// lifecycle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#lifecycle AwsBackupPlan#lifecycle}
	// Experimental.
	Lifecycle *AwsBackupPlan_RuleLifecycleProperty `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#recovery_point_tags AwsBackupPlan#recovery_point_tags}.
	// Experimental.
	RecoveryPointTags *map[string]*string `field:"optional" json:"recoveryPointTags" yaml:"recoveryPointTags"`
	// scan_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#scan_action AwsBackupPlan#scan_action}
	// Experimental.
	ScanAction interface{} `field:"optional" json:"scanAction" yaml:"scanAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#schedule AwsBackupPlan#schedule}.
	// Experimental.
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#schedule_expression_timezone AwsBackupPlan#schedule_expression_timezone}.
	// Experimental.
	ScheduleExpressionTimezone *string `field:"optional" json:"scheduleExpressionTimezone" yaml:"scheduleExpressionTimezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#start_window AwsBackupPlan#start_window}.
	// Experimental.
	StartWindow *float64 `field:"optional" json:"startWindow" yaml:"startWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#target_logically_air_gapped_backup_vault_arn AwsBackupPlan#target_logically_air_gapped_backup_vault_arn}.
	// Experimental.
	TargetLogicallyAirGappedBackupVaultArn *string `field:"optional" json:"targetLogicallyAirGappedBackupVaultArn" yaml:"targetLogicallyAirGappedBackupVaultArn"`
}

