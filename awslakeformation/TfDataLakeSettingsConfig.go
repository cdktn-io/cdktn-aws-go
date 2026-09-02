package awslakeformation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataLakeSettingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#admins TfDataLakeSettings#admins}.
	// Experimental.
	Admins *[]*string `field:"optional" json:"admins" yaml:"admins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#allow_external_data_filtering TfDataLakeSettings#allow_external_data_filtering}.
	// Experimental.
	AllowExternalDataFiltering interface{} `field:"optional" json:"allowExternalDataFiltering" yaml:"allowExternalDataFiltering"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#allow_full_table_external_data_access TfDataLakeSettings#allow_full_table_external_data_access}.
	// Experimental.
	AllowFullTableExternalDataAccess interface{} `field:"optional" json:"allowFullTableExternalDataAccess" yaml:"allowFullTableExternalDataAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#authorized_session_tag_value_list TfDataLakeSettings#authorized_session_tag_value_list}.
	// Experimental.
	AuthorizedSessionTagValueList *[]*string `field:"optional" json:"authorizedSessionTagValueList" yaml:"authorizedSessionTagValueList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#catalog_id TfDataLakeSettings#catalog_id}.
	// Experimental.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// create_database_default_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#create_database_default_permissions TfDataLakeSettings#create_database_default_permissions}
	// Experimental.
	CreateDatabaseDefaultPermissions interface{} `field:"optional" json:"createDatabaseDefaultPermissions" yaml:"createDatabaseDefaultPermissions"`
	// create_table_default_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#create_table_default_permissions TfDataLakeSettings#create_table_default_permissions}
	// Experimental.
	CreateTableDefaultPermissions interface{} `field:"optional" json:"createTableDefaultPermissions" yaml:"createTableDefaultPermissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#external_data_filtering_allow_list TfDataLakeSettings#external_data_filtering_allow_list}.
	// Experimental.
	ExternalDataFilteringAllowList *[]*string `field:"optional" json:"externalDataFilteringAllowList" yaml:"externalDataFilteringAllowList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#id TfDataLakeSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#parameters TfDataLakeSettings#parameters}.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#read_only_admins TfDataLakeSettings#read_only_admins}.
	// Experimental.
	ReadOnlyAdmins *[]*string `field:"optional" json:"readOnlyAdmins" yaml:"readOnlyAdmins"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#region TfDataLakeSettings#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#trusted_resource_owners TfDataLakeSettings#trusted_resource_owners}.
	// Experimental.
	TrustedResourceOwners *[]*string `field:"optional" json:"trustedResourceOwners" yaml:"trustedResourceOwners"`
}

