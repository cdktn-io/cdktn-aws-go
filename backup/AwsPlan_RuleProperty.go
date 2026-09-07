package backup


// Experimental.
type AwsPlan_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#rule_name AwsPlan#rule_name}.
	// Experimental.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#target_vault_name AwsPlan#target_vault_name}.
	// Experimental.
	TargetVaultName *string `field:"required" json:"targetVaultName" yaml:"targetVaultName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#completion_window AwsPlan#completion_window}.
	// Experimental.
	CompletionWindow *float64 `field:"optional" json:"completionWindow" yaml:"completionWindow"`
	// copy_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#copy_action AwsPlan#copy_action}
	// Experimental.
	CopyAction interface{} `field:"optional" json:"copyAction" yaml:"copyAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#enable_continuous_backup AwsPlan#enable_continuous_backup}.
	// Experimental.
	EnableContinuousBackup interface{} `field:"optional" json:"enableContinuousBackup" yaml:"enableContinuousBackup"`
	// lifecycle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#lifecycle AwsPlan#lifecycle}
	// Experimental.
	Lifecycle *AwsPlan_RuleLifecycleProperty `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#recovery_point_tags AwsPlan#recovery_point_tags}.
	// Experimental.
	RecoveryPointTags *map[string]*string `field:"optional" json:"recoveryPointTags" yaml:"recoveryPointTags"`
	// scan_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#scan_action AwsPlan#scan_action}
	// Experimental.
	ScanAction interface{} `field:"optional" json:"scanAction" yaml:"scanAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#schedule AwsPlan#schedule}.
	// Experimental.
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#schedule_expression_timezone AwsPlan#schedule_expression_timezone}.
	// Experimental.
	ScheduleExpressionTimezone *string `field:"optional" json:"scheduleExpressionTimezone" yaml:"scheduleExpressionTimezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#start_window AwsPlan#start_window}.
	// Experimental.
	StartWindow *float64 `field:"optional" json:"startWindow" yaml:"startWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#target_logically_air_gapped_backup_vault_arn AwsPlan#target_logically_air_gapped_backup_vault_arn}.
	// Experimental.
	TargetLogicallyAirGappedBackupVaultArn *string `field:"optional" json:"targetLogicallyAirGappedBackupVaultArn" yaml:"targetLogicallyAirGappedBackupVaultArn"`
}

