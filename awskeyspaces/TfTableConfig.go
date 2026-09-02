package awskeyspaces

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTableConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#keyspace_name TfTable#keyspace_name}.
	// Experimental.
	KeyspaceName *string `field:"required" json:"keyspaceName" yaml:"keyspaceName"`
	// schema_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#schema_definition TfTable#schema_definition}
	// Experimental.
	SchemaDefinition *TfTable_SchemaDefinitionProperty `field:"required" json:"schemaDefinition" yaml:"schemaDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#table_name TfTable#table_name}.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// capacity_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#capacity_specification TfTable#capacity_specification}
	// Experimental.
	CapacitySpecification *TfTable_CapacitySpecificationProperty `field:"optional" json:"capacitySpecification" yaml:"capacitySpecification"`
	// client_side_timestamps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#client_side_timestamps TfTable#client_side_timestamps}
	// Experimental.
	ClientSideTimestamps *TfTable_ClientSideTimestampsProperty `field:"optional" json:"clientSideTimestamps" yaml:"clientSideTimestamps"`
	// comment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#comment TfTable#comment}
	// Experimental.
	Comment *TfTable_CommentProperty `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#default_time_to_live TfTable#default_time_to_live}.
	// Experimental.
	DefaultTimeToLive *float64 `field:"optional" json:"defaultTimeToLive" yaml:"defaultTimeToLive"`
	// encryption_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#encryption_specification TfTable#encryption_specification}
	// Experimental.
	EncryptionSpecification *TfTable_EncryptionSpecificationProperty `field:"optional" json:"encryptionSpecification" yaml:"encryptionSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#id TfTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// point_in_time_recovery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#point_in_time_recovery TfTable#point_in_time_recovery}
	// Experimental.
	PointInTimeRecovery *TfTable_PointInTimeRecoveryProperty `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#region TfTable#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#tags TfTable#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#tags_all TfTable#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#timeouts TfTable#timeouts}
	// Experimental.
	Timeouts *TfTable_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#ttl TfTable#ttl}
	// Experimental.
	Ttl *TfTable_TtlProperty `field:"optional" json:"ttl" yaml:"ttl"`
}

