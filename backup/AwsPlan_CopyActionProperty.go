package backup


// Experimental.
type AwsPlan_CopyActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#destination_vault_arn AwsPlan#destination_vault_arn}.
	// Experimental.
	DestinationVaultArn *string `field:"required" json:"destinationVaultArn" yaml:"destinationVaultArn"`
	// lifecycle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#lifecycle AwsPlan#lifecycle}
	// Experimental.
	Lifecycle *AwsPlan_RuleCopyActionLifecycleProperty `field:"optional" json:"lifecycle" yaml:"lifecycle"`
}

