package backup


// Experimental.
type AwsRestoreTestingPlan_RecoveryPointSelectionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_plan#algorithm AwsRestoreTestingPlan#algorithm}.
	// Experimental.
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_plan#include_vaults AwsRestoreTestingPlan#include_vaults}.
	// Experimental.
	IncludeVaults *[]*string `field:"required" json:"includeVaults" yaml:"includeVaults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_plan#recovery_point_types AwsRestoreTestingPlan#recovery_point_types}.
	// Experimental.
	RecoveryPointTypes *[]*string `field:"required" json:"recoveryPointTypes" yaml:"recoveryPointTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_plan#exclude_vaults AwsRestoreTestingPlan#exclude_vaults}.
	// Experimental.
	ExcludeVaults *[]*string `field:"optional" json:"excludeVaults" yaml:"excludeVaults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_plan#selection_window_days AwsRestoreTestingPlan#selection_window_days}.
	// Experimental.
	SelectionWindowDays *float64 `field:"optional" json:"selectionWindowDays" yaml:"selectionWindowDays"`
}

