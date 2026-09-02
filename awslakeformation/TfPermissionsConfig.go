package awslakeformation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPermissionsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#permissions TfPermissions#permissions}.
	// Experimental.
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#principal TfPermissions#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#catalog_id TfPermissions#catalog_id}.
	// Experimental.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#catalog_resource TfPermissions#catalog_resource}.
	// Experimental.
	CatalogResource interface{} `field:"optional" json:"catalogResource" yaml:"catalogResource"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#database TfPermissions#database}
	// Experimental.
	Database *TfPermissions_DatabaseProperty `field:"optional" json:"database" yaml:"database"`
	// data_cells_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#data_cells_filter TfPermissions#data_cells_filter}
	// Experimental.
	DataCellsFilter *TfPermissions_DataCellsFilterProperty `field:"optional" json:"dataCellsFilter" yaml:"dataCellsFilter"`
	// data_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#data_location TfPermissions#data_location}
	// Experimental.
	DataLocation *TfPermissions_DataLocationProperty `field:"optional" json:"dataLocation" yaml:"dataLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#id TfPermissions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lf_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#lf_tag TfPermissions#lf_tag}
	// Experimental.
	LfTag *TfPermissions_LfTagProperty `field:"optional" json:"lfTag" yaml:"lfTag"`
	// lf_tag_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#lf_tag_policy TfPermissions#lf_tag_policy}
	// Experimental.
	LfTagPolicy *TfPermissions_LfTagPolicyProperty `field:"optional" json:"lfTagPolicy" yaml:"lfTagPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#permissions_with_grant_option TfPermissions#permissions_with_grant_option}.
	// Experimental.
	PermissionsWithGrantOption *[]*string `field:"optional" json:"permissionsWithGrantOption" yaml:"permissionsWithGrantOption"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#region TfPermissions#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#table TfPermissions#table}
	// Experimental.
	Table *TfPermissions_TableProperty `field:"optional" json:"table" yaml:"table"`
	// table_with_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#table_with_columns TfPermissions#table_with_columns}
	// Experimental.
	TableWithColumns *TfPermissions_TableWithColumnsProperty `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
}

