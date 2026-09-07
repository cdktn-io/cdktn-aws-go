package lakeformation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsResourceLfTagsConfig struct {
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
	// lf_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_resource_lf_tags#lf_tag AwsResourceLfTags#lf_tag}
	// Experimental.
	LfTag interface{} `field:"required" json:"lfTag" yaml:"lfTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_resource_lf_tags#catalog_id AwsResourceLfTags#catalog_id}.
	// Experimental.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_resource_lf_tags#database AwsResourceLfTags#database}
	// Experimental.
	Database *AwsResourceLfTags_DatabaseProperty `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_resource_lf_tags#id AwsResourceLfTags#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_resource_lf_tags#region AwsResourceLfTags#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_resource_lf_tags#table AwsResourceLfTags#table}
	// Experimental.
	Table *AwsResourceLfTags_TableProperty `field:"optional" json:"table" yaml:"table"`
	// table_with_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_resource_lf_tags#table_with_columns AwsResourceLfTags#table_with_columns}
	// Experimental.
	TableWithColumns *AwsResourceLfTags_TableWithColumnsProperty `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_resource_lf_tags#timeouts AwsResourceLfTags#timeouts}
	// Experimental.
	Timeouts *AwsResourceLfTags_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

