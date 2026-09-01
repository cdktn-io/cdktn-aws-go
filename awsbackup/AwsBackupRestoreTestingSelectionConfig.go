package awsbackup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBackupRestoreTestingSelectionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_selection#iam_role_arn AwsBackupRestoreTestingSelection#iam_role_arn}.
	// Experimental.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_selection#name AwsBackupRestoreTestingSelection#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_selection#protected_resource_type AwsBackupRestoreTestingSelection#protected_resource_type}.
	// Experimental.
	ProtectedResourceType *string `field:"required" json:"protectedResourceType" yaml:"protectedResourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_selection#restore_testing_plan_name AwsBackupRestoreTestingSelection#restore_testing_plan_name}.
	// Experimental.
	RestoreTestingPlanName *string `field:"required" json:"restoreTestingPlanName" yaml:"restoreTestingPlanName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_selection#protected_resource_arns AwsBackupRestoreTestingSelection#protected_resource_arns}.
	// Experimental.
	ProtectedResourceArns *[]*string `field:"optional" json:"protectedResourceArns" yaml:"protectedResourceArns"`
	// protected_resource_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_selection#protected_resource_conditions AwsBackupRestoreTestingSelection#protected_resource_conditions}
	// Experimental.
	ProtectedResourceConditions interface{} `field:"optional" json:"protectedResourceConditions" yaml:"protectedResourceConditions"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_selection#region AwsBackupRestoreTestingSelection#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_selection#restore_metadata_overrides AwsBackupRestoreTestingSelection#restore_metadata_overrides}.
	// Experimental.
	RestoreMetadataOverrides *map[string]*string `field:"optional" json:"restoreMetadataOverrides" yaml:"restoreMetadataOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_selection#validation_window_hours AwsBackupRestoreTestingSelection#validation_window_hours}.
	// Experimental.
	ValidationWindowHours *float64 `field:"optional" json:"validationWindowHours" yaml:"validationWindowHours"`
}

