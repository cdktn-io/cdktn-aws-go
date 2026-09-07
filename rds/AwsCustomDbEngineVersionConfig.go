package rds

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCustomDbEngineVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#engine AwsCustomDbEngineVersion#engine}.
	// Experimental.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#engine_version AwsCustomDbEngineVersion#engine_version}.
	// Experimental.
	EngineVersion *string `field:"required" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#database_installation_files_s3_bucket_name AwsCustomDbEngineVersion#database_installation_files_s3_bucket_name}.
	// Experimental.
	DatabaseInstallationFilesS3BucketName *string `field:"optional" json:"databaseInstallationFilesS3BucketName" yaml:"databaseInstallationFilesS3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#database_installation_files_s3_prefix AwsCustomDbEngineVersion#database_installation_files_s3_prefix}.
	// Experimental.
	DatabaseInstallationFilesS3Prefix *string `field:"optional" json:"databaseInstallationFilesS3Prefix" yaml:"databaseInstallationFilesS3Prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#description AwsCustomDbEngineVersion#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#filename AwsCustomDbEngineVersion#filename}.
	// Experimental.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#id AwsCustomDbEngineVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#kms_key_id AwsCustomDbEngineVersion#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#manifest AwsCustomDbEngineVersion#manifest}.
	// Experimental.
	Manifest *string `field:"optional" json:"manifest" yaml:"manifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#manifest_hash AwsCustomDbEngineVersion#manifest_hash}.
	// Experimental.
	ManifestHash *string `field:"optional" json:"manifestHash" yaml:"manifestHash"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#region AwsCustomDbEngineVersion#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#source_image_id AwsCustomDbEngineVersion#source_image_id}.
	// Experimental.
	SourceImageId *string `field:"optional" json:"sourceImageId" yaml:"sourceImageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#status AwsCustomDbEngineVersion#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#tags AwsCustomDbEngineVersion#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#tags_all AwsCustomDbEngineVersion#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#timeouts AwsCustomDbEngineVersion#timeouts}
	// Experimental.
	Timeouts *AwsCustomDbEngineVersion_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

