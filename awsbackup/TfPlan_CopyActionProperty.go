package awsbackup


// Experimental.
type TfPlan_CopyActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#destination_vault_arn TfPlan#destination_vault_arn}.
	// Experimental.
	DestinationVaultArn *string `field:"required" json:"destinationVaultArn" yaml:"destinationVaultArn"`
	// lifecycle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#lifecycle TfPlan#lifecycle}
	// Experimental.
	Lifecycle *TfPlan_RuleCopyActionLifecycleProperty `field:"optional" json:"lifecycle" yaml:"lifecycle"`
}

