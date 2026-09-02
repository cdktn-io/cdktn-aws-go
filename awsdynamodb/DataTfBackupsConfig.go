package awsdynamodb

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfBackupsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/dynamodb_backups#backup_type DataTfBackups#backup_type}.
	// Experimental.
	BackupType *string `field:"optional" json:"backupType" yaml:"backupType"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/dynamodb_backups#region DataTfBackups#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/dynamodb_backups#table_name DataTfBackups#table_name}.
	// Experimental.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/dynamodb_backups#time_range_lower_bound DataTfBackups#time_range_lower_bound}.
	// Experimental.
	TimeRangeLowerBound *string `field:"optional" json:"timeRangeLowerBound" yaml:"timeRangeLowerBound"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/dynamodb_backups#time_range_upper_bound DataTfBackups#time_range_upper_bound}.
	// Experimental.
	TimeRangeUpperBound *string `field:"optional" json:"timeRangeUpperBound" yaml:"timeRangeUpperBound"`
}

