package awsstoragegateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfNfsFileShareConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#client_list TfNfsFileShare#client_list}.
	// Experimental.
	ClientList *[]*string `field:"required" json:"clientList" yaml:"clientList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#gateway_arn TfNfsFileShare#gateway_arn}.
	// Experimental.
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#location_arn TfNfsFileShare#location_arn}.
	// Experimental.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#role_arn TfNfsFileShare#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#audit_destination_arn TfNfsFileShare#audit_destination_arn}.
	// Experimental.
	AuditDestinationArn *string `field:"optional" json:"auditDestinationArn" yaml:"auditDestinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#bucket_region TfNfsFileShare#bucket_region}.
	// Experimental.
	BucketRegion *string `field:"optional" json:"bucketRegion" yaml:"bucketRegion"`
	// cache_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#cache_attributes TfNfsFileShare#cache_attributes}
	// Experimental.
	CacheAttributes *TfNfsFileShare_CacheAttributesProperty `field:"optional" json:"cacheAttributes" yaml:"cacheAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#default_storage_class TfNfsFileShare#default_storage_class}.
	// Experimental.
	DefaultStorageClass *string `field:"optional" json:"defaultStorageClass" yaml:"defaultStorageClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#file_share_name TfNfsFileShare#file_share_name}.
	// Experimental.
	FileShareName *string `field:"optional" json:"fileShareName" yaml:"fileShareName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#guess_mime_type_enabled TfNfsFileShare#guess_mime_type_enabled}.
	// Experimental.
	GuessMimeTypeEnabled interface{} `field:"optional" json:"guessMimeTypeEnabled" yaml:"guessMimeTypeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#id TfNfsFileShare#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#kms_encrypted TfNfsFileShare#kms_encrypted}.
	// Experimental.
	KmsEncrypted interface{} `field:"optional" json:"kmsEncrypted" yaml:"kmsEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#kms_key_arn TfNfsFileShare#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// nfs_file_share_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#nfs_file_share_defaults TfNfsFileShare#nfs_file_share_defaults}
	// Experimental.
	NfsFileShareDefaults *TfNfsFileShare_NfsFileShareDefaultsProperty `field:"optional" json:"nfsFileShareDefaults" yaml:"nfsFileShareDefaults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#notification_policy TfNfsFileShare#notification_policy}.
	// Experimental.
	NotificationPolicy *string `field:"optional" json:"notificationPolicy" yaml:"notificationPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#object_acl TfNfsFileShare#object_acl}.
	// Experimental.
	ObjectAcl *string `field:"optional" json:"objectAcl" yaml:"objectAcl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#read_only TfNfsFileShare#read_only}.
	// Experimental.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#region TfNfsFileShare#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#requester_pays TfNfsFileShare#requester_pays}.
	// Experimental.
	RequesterPays interface{} `field:"optional" json:"requesterPays" yaml:"requesterPays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#squash TfNfsFileShare#squash}.
	// Experimental.
	Squash *string `field:"optional" json:"squash" yaml:"squash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#tags TfNfsFileShare#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#tags_all TfNfsFileShare#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#timeouts TfNfsFileShare#timeouts}
	// Experimental.
	Timeouts *TfNfsFileShare_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#vpc_endpoint_dns_name TfNfsFileShare#vpc_endpoint_dns_name}.
	// Experimental.
	VpcEndpointDnsName *string `field:"optional" json:"vpcEndpointDnsName" yaml:"vpcEndpointDnsName"`
}

