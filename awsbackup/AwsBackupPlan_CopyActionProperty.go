package awsbackup


// Experimental.
type AwsBackupPlan_CopyActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#destination_vault_arn AwsBackupPlan#destination_vault_arn}.
	// Experimental.
	DestinationVaultArn *string `field:"required" json:"destinationVaultArn" yaml:"destinationVaultArn"`
	// lifecycle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#lifecycle AwsBackupPlan#lifecycle}
	// Experimental.
	Lifecycle *AwsBackupPlan_RuleCopyActionLifecycleProperty `field:"optional" json:"lifecycle" yaml:"lifecycle"`
}

