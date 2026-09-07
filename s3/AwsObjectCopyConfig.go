package s3

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsObjectCopyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#bucket AwsObjectCopy#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#key AwsObjectCopy#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#source AwsObjectCopy#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#acl AwsObjectCopy#acl}.
	// Experimental.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#bucket_key_enabled AwsObjectCopy#bucket_key_enabled}.
	// Experimental.
	BucketKeyEnabled interface{} `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#cache_control AwsObjectCopy#cache_control}.
	// Experimental.
	CacheControl *string `field:"optional" json:"cacheControl" yaml:"cacheControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#checksum_algorithm AwsObjectCopy#checksum_algorithm}.
	// Experimental.
	ChecksumAlgorithm *string `field:"optional" json:"checksumAlgorithm" yaml:"checksumAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#content_disposition AwsObjectCopy#content_disposition}.
	// Experimental.
	ContentDisposition *string `field:"optional" json:"contentDisposition" yaml:"contentDisposition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#content_encoding AwsObjectCopy#content_encoding}.
	// Experimental.
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#content_language AwsObjectCopy#content_language}.
	// Experimental.
	ContentLanguage *string `field:"optional" json:"contentLanguage" yaml:"contentLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#content_type AwsObjectCopy#content_type}.
	// Experimental.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#copy_if_match AwsObjectCopy#copy_if_match}.
	// Experimental.
	CopyIfMatch *string `field:"optional" json:"copyIfMatch" yaml:"copyIfMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#copy_if_modified_since AwsObjectCopy#copy_if_modified_since}.
	// Experimental.
	CopyIfModifiedSince *string `field:"optional" json:"copyIfModifiedSince" yaml:"copyIfModifiedSince"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#copy_if_none_match AwsObjectCopy#copy_if_none_match}.
	// Experimental.
	CopyIfNoneMatch *string `field:"optional" json:"copyIfNoneMatch" yaml:"copyIfNoneMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#copy_if_unmodified_since AwsObjectCopy#copy_if_unmodified_since}.
	// Experimental.
	CopyIfUnmodifiedSince *string `field:"optional" json:"copyIfUnmodifiedSince" yaml:"copyIfUnmodifiedSince"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#customer_algorithm AwsObjectCopy#customer_algorithm}.
	// Experimental.
	CustomerAlgorithm *string `field:"optional" json:"customerAlgorithm" yaml:"customerAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#customer_key AwsObjectCopy#customer_key}.
	// Experimental.
	CustomerKey *string `field:"optional" json:"customerKey" yaml:"customerKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#customer_key_md5 AwsObjectCopy#customer_key_md5}.
	// Experimental.
	CustomerKeyMd5 *string `field:"optional" json:"customerKeyMd5" yaml:"customerKeyMd5"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#expected_bucket_owner AwsObjectCopy#expected_bucket_owner}.
	// Experimental.
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#expected_source_bucket_owner AwsObjectCopy#expected_source_bucket_owner}.
	// Experimental.
	ExpectedSourceBucketOwner *string `field:"optional" json:"expectedSourceBucketOwner" yaml:"expectedSourceBucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#expires AwsObjectCopy#expires}.
	// Experimental.
	Expires *string `field:"optional" json:"expires" yaml:"expires"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#force_destroy AwsObjectCopy#force_destroy}.
	// Experimental.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// grant block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#grant AwsObjectCopy#grant}
	// Experimental.
	Grant interface{} `field:"optional" json:"grant" yaml:"grant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#id AwsObjectCopy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#kms_encryption_context AwsObjectCopy#kms_encryption_context}.
	// Experimental.
	KmsEncryptionContext *string `field:"optional" json:"kmsEncryptionContext" yaml:"kmsEncryptionContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#kms_key_id AwsObjectCopy#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#metadata AwsObjectCopy#metadata}.
	// Experimental.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#metadata_directive AwsObjectCopy#metadata_directive}.
	// Experimental.
	MetadataDirective *string `field:"optional" json:"metadataDirective" yaml:"metadataDirective"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#object_lock_legal_hold_status AwsObjectCopy#object_lock_legal_hold_status}.
	// Experimental.
	ObjectLockLegalHoldStatus *string `field:"optional" json:"objectLockLegalHoldStatus" yaml:"objectLockLegalHoldStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#object_lock_mode AwsObjectCopy#object_lock_mode}.
	// Experimental.
	ObjectLockMode *string `field:"optional" json:"objectLockMode" yaml:"objectLockMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#object_lock_retain_until_date AwsObjectCopy#object_lock_retain_until_date}.
	// Experimental.
	ObjectLockRetainUntilDate *string `field:"optional" json:"objectLockRetainUntilDate" yaml:"objectLockRetainUntilDate"`
	// override_provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#override_provider AwsObjectCopy#override_provider}
	// Experimental.
	OverrideProvider *AwsObjectCopy_OverrideProviderProperty `field:"optional" json:"overrideProvider" yaml:"overrideProvider"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#region AwsObjectCopy#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#request_payer AwsObjectCopy#request_payer}.
	// Experimental.
	RequestPayer *string `field:"optional" json:"requestPayer" yaml:"requestPayer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#server_side_encryption AwsObjectCopy#server_side_encryption}.
	// Experimental.
	ServerSideEncryption *string `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#source_customer_algorithm AwsObjectCopy#source_customer_algorithm}.
	// Experimental.
	SourceCustomerAlgorithm *string `field:"optional" json:"sourceCustomerAlgorithm" yaml:"sourceCustomerAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#source_customer_key AwsObjectCopy#source_customer_key}.
	// Experimental.
	SourceCustomerKey *string `field:"optional" json:"sourceCustomerKey" yaml:"sourceCustomerKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#source_customer_key_md5 AwsObjectCopy#source_customer_key_md5}.
	// Experimental.
	SourceCustomerKeyMd5 *string `field:"optional" json:"sourceCustomerKeyMd5" yaml:"sourceCustomerKeyMd5"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#storage_class AwsObjectCopy#storage_class}.
	// Experimental.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#tagging_directive AwsObjectCopy#tagging_directive}.
	// Experimental.
	TaggingDirective *string `field:"optional" json:"taggingDirective" yaml:"taggingDirective"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#tags AwsObjectCopy#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#tags_all AwsObjectCopy#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#website_redirect AwsObjectCopy#website_redirect}.
	// Experimental.
	WebsiteRedirect *string `field:"optional" json:"websiteRedirect" yaml:"websiteRedirect"`
}

