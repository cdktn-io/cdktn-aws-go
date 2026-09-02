package awsquicksight

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#data_set_id TfDataSet#data_set_id}.
	// Experimental.
	DataSetId *string `field:"required" json:"dataSetId" yaml:"dataSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#import_mode TfDataSet#import_mode}.
	// Experimental.
	ImportMode *string `field:"required" json:"importMode" yaml:"importMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#name TfDataSet#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#aws_account_id TfDataSet#aws_account_id}.
	// Experimental.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// column_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_groups TfDataSet#column_groups}
	// Experimental.
	ColumnGroups interface{} `field:"optional" json:"columnGroups" yaml:"columnGroups"`
	// column_level_permission_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_level_permission_rules TfDataSet#column_level_permission_rules}
	// Experimental.
	ColumnLevelPermissionRules interface{} `field:"optional" json:"columnLevelPermissionRules" yaml:"columnLevelPermissionRules"`
	// data_set_usage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#data_set_usage_configuration TfDataSet#data_set_usage_configuration}
	// Experimental.
	DataSetUsageConfiguration *TfDataSet_DataSetUsageConfigurationProperty `field:"optional" json:"dataSetUsageConfiguration" yaml:"dataSetUsageConfiguration"`
	// field_folders block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#field_folders TfDataSet#field_folders}
	// Experimental.
	FieldFolders interface{} `field:"optional" json:"fieldFolders" yaml:"fieldFolders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#id TfDataSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// logical_table_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#logical_table_map TfDataSet#logical_table_map}
	// Experimental.
	LogicalTableMap interface{} `field:"optional" json:"logicalTableMap" yaml:"logicalTableMap"`
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#permissions TfDataSet#permissions}
	// Experimental.
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// physical_table_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#physical_table_map TfDataSet#physical_table_map}
	// Experimental.
	PhysicalTableMap interface{} `field:"optional" json:"physicalTableMap" yaml:"physicalTableMap"`
	// refresh_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#refresh_properties TfDataSet#refresh_properties}
	// Experimental.
	RefreshProperties *TfDataSet_RefreshPropertiesProperty `field:"optional" json:"refreshProperties" yaml:"refreshProperties"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#region TfDataSet#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// row_level_permission_data_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#row_level_permission_data_set TfDataSet#row_level_permission_data_set}
	// Experimental.
	RowLevelPermissionDataSet *TfDataSet_RowLevelPermissionDataSetProperty `field:"optional" json:"rowLevelPermissionDataSet" yaml:"rowLevelPermissionDataSet"`
	// row_level_permission_tag_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#row_level_permission_tag_configuration TfDataSet#row_level_permission_tag_configuration}
	// Experimental.
	RowLevelPermissionTagConfiguration *TfDataSet_RowLevelPermissionTagConfigurationProperty `field:"optional" json:"rowLevelPermissionTagConfiguration" yaml:"rowLevelPermissionTagConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#tags TfDataSet#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#tags_all TfDataSet#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#use_as TfDataSet#use_as}.
	// Experimental.
	UseAs *string `field:"optional" json:"useAs" yaml:"useAs"`
}

