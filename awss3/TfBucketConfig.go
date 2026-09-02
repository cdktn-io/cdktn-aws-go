package awss3

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfBucketConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#acceleration_status TfBucket#acceleration_status}.
	// Experimental.
	AccelerationStatus *string `field:"optional" json:"accelerationStatus" yaml:"accelerationStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#acl TfBucket#acl}.
	// Experimental.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#bucket TfBucket#bucket}.
	// Experimental.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#bucket_namespace TfBucket#bucket_namespace}.
	// Experimental.
	BucketNamespace *string `field:"optional" json:"bucketNamespace" yaml:"bucketNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#bucket_prefix TfBucket#bucket_prefix}.
	// Experimental.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// cors_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#cors_rule TfBucket#cors_rule}
	// Experimental.
	CorsRule interface{} `field:"optional" json:"corsRule" yaml:"corsRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#force_destroy TfBucket#force_destroy}.
	// Experimental.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// grant block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#grant TfBucket#grant}
	// Experimental.
	Grant interface{} `field:"optional" json:"grant" yaml:"grant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#id TfBucket#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lifecycle_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#lifecycle_rule TfBucket#lifecycle_rule}
	// Experimental.
	LifecycleRule interface{} `field:"optional" json:"lifecycleRule" yaml:"lifecycleRule"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#logging TfBucket#logging}
	// Experimental.
	Logging *TfBucket_LoggingProperty `field:"optional" json:"logging" yaml:"logging"`
	// object_lock_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#object_lock_configuration TfBucket#object_lock_configuration}
	// Experimental.
	ObjectLockConfiguration *TfBucket_ObjectLockConfigurationProperty `field:"optional" json:"objectLockConfiguration" yaml:"objectLockConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#object_lock_enabled TfBucket#object_lock_enabled}.
	// Experimental.
	ObjectLockEnabled interface{} `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#policy TfBucket#policy}.
	// Experimental.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#region TfBucket#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// replication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#replication_configuration TfBucket#replication_configuration}
	// Experimental.
	ReplicationConfiguration *TfBucket_ReplicationConfigurationProperty `field:"optional" json:"replicationConfiguration" yaml:"replicationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#request_payer TfBucket#request_payer}.
	// Experimental.
	RequestPayer *string `field:"optional" json:"requestPayer" yaml:"requestPayer"`
	// server_side_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#server_side_encryption_configuration TfBucket#server_side_encryption_configuration}
	// Experimental.
	ServerSideEncryptionConfiguration *TfBucket_ServerSideEncryptionConfigurationProperty `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#tags TfBucket#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#tags_all TfBucket#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#timeouts TfBucket#timeouts}
	// Experimental.
	Timeouts *TfBucket_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// versioning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#versioning TfBucket#versioning}
	// Experimental.
	Versioning *TfBucket_VersioningProperty `field:"optional" json:"versioning" yaml:"versioning"`
	// website block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#website TfBucket#website}
	// Experimental.
	Website *TfBucket_WebsiteProperty `field:"optional" json:"website" yaml:"website"`
}

