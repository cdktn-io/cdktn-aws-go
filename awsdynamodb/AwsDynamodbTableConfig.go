package awsdynamodb

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDynamodbTableConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#name AwsDynamodbTable#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// attribute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#attribute AwsDynamodbTable#attribute}
	// Experimental.
	Attribute interface{} `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#billing_mode AwsDynamodbTable#billing_mode}.
	// Experimental.
	BillingMode *string `field:"optional" json:"billingMode" yaml:"billingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#deletion_protection_enabled AwsDynamodbTable#deletion_protection_enabled}.
	// Experimental.
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// global_secondary_index block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#global_secondary_index AwsDynamodbTable#global_secondary_index}
	// Experimental.
	GlobalSecondaryIndex interface{} `field:"optional" json:"globalSecondaryIndex" yaml:"globalSecondaryIndex"`
	// global_table_witness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#global_table_witness AwsDynamodbTable#global_table_witness}
	// Experimental.
	GlobalTableWitness *AwsDynamodbTable_GlobalTableWitnessProperty `field:"optional" json:"globalTableWitness" yaml:"globalTableWitness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#hash_key AwsDynamodbTable#hash_key}.
	// Experimental.
	HashKey *string `field:"optional" json:"hashKey" yaml:"hashKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#id AwsDynamodbTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// import_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#import_table AwsDynamodbTable#import_table}
	// Experimental.
	ImportTable *AwsDynamodbTable_ImportTableProperty `field:"optional" json:"importTable" yaml:"importTable"`
	// local_secondary_index block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#local_secondary_index AwsDynamodbTable#local_secondary_index}
	// Experimental.
	LocalSecondaryIndex interface{} `field:"optional" json:"localSecondaryIndex" yaml:"localSecondaryIndex"`
	// on_demand_throughput block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#on_demand_throughput AwsDynamodbTable#on_demand_throughput}
	// Experimental.
	OnDemandThroughput *AwsDynamodbTable_OnDemandThroughputProperty `field:"optional" json:"onDemandThroughput" yaml:"onDemandThroughput"`
	// point_in_time_recovery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#point_in_time_recovery AwsDynamodbTable#point_in_time_recovery}
	// Experimental.
	PointInTimeRecovery *AwsDynamodbTable_PointInTimeRecoveryProperty `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#range_key AwsDynamodbTable#range_key}.
	// Experimental.
	RangeKey *string `field:"optional" json:"rangeKey" yaml:"rangeKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#read_capacity AwsDynamodbTable#read_capacity}.
	// Experimental.
	ReadCapacity *float64 `field:"optional" json:"readCapacity" yaml:"readCapacity"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#region AwsDynamodbTable#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// replica block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#replica AwsDynamodbTable#replica}
	// Experimental.
	Replica interface{} `field:"optional" json:"replica" yaml:"replica"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#restore_backup_arn AwsDynamodbTable#restore_backup_arn}.
	// Experimental.
	RestoreBackupArn *string `field:"optional" json:"restoreBackupArn" yaml:"restoreBackupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#restore_date_time AwsDynamodbTable#restore_date_time}.
	// Experimental.
	RestoreDateTime *string `field:"optional" json:"restoreDateTime" yaml:"restoreDateTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#restore_source_name AwsDynamodbTable#restore_source_name}.
	// Experimental.
	RestoreSourceName *string `field:"optional" json:"restoreSourceName" yaml:"restoreSourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#restore_source_table_arn AwsDynamodbTable#restore_source_table_arn}.
	// Experimental.
	RestoreSourceTableArn *string `field:"optional" json:"restoreSourceTableArn" yaml:"restoreSourceTableArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#restore_to_latest_time AwsDynamodbTable#restore_to_latest_time}.
	// Experimental.
	RestoreToLatestTime interface{} `field:"optional" json:"restoreToLatestTime" yaml:"restoreToLatestTime"`
	// server_side_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#server_side_encryption AwsDynamodbTable#server_side_encryption}
	// Experimental.
	ServerSideEncryption *AwsDynamodbTable_ServerSideEncryptionProperty `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#stream_enabled AwsDynamodbTable#stream_enabled}.
	// Experimental.
	StreamEnabled interface{} `field:"optional" json:"streamEnabled" yaml:"streamEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#stream_view_type AwsDynamodbTable#stream_view_type}.
	// Experimental.
	StreamViewType *string `field:"optional" json:"streamViewType" yaml:"streamViewType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#table_class AwsDynamodbTable#table_class}.
	// Experimental.
	TableClass *string `field:"optional" json:"tableClass" yaml:"tableClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#tags AwsDynamodbTable#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#tags_all AwsDynamodbTable#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#timeouts AwsDynamodbTable#timeouts}
	// Experimental.
	Timeouts *AwsDynamodbTable_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#ttl AwsDynamodbTable#ttl}
	// Experimental.
	Ttl *AwsDynamodbTable_TtlProperty `field:"optional" json:"ttl" yaml:"ttl"`
	// warm_throughput block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#warm_throughput AwsDynamodbTable#warm_throughput}
	// Experimental.
	WarmThroughput *AwsDynamodbTable_WarmThroughputProperty `field:"optional" json:"warmThroughput" yaml:"warmThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#write_capacity AwsDynamodbTable#write_capacity}.
	// Experimental.
	WriteCapacity *float64 `field:"optional" json:"writeCapacity" yaml:"writeCapacity"`
}

