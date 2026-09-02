package awsstoragegateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSmbFileShareConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#gateway_arn TfSmbFileShare#gateway_arn}.
	// Experimental.
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#location_arn TfSmbFileShare#location_arn}.
	// Experimental.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#role_arn TfSmbFileShare#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#access_based_enumeration TfSmbFileShare#access_based_enumeration}.
	// Experimental.
	AccessBasedEnumeration interface{} `field:"optional" json:"accessBasedEnumeration" yaml:"accessBasedEnumeration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#admin_user_list TfSmbFileShare#admin_user_list}.
	// Experimental.
	AdminUserList *[]*string `field:"optional" json:"adminUserList" yaml:"adminUserList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#audit_destination_arn TfSmbFileShare#audit_destination_arn}.
	// Experimental.
	AuditDestinationArn *string `field:"optional" json:"auditDestinationArn" yaml:"auditDestinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#authentication TfSmbFileShare#authentication}.
	// Experimental.
	Authentication *string `field:"optional" json:"authentication" yaml:"authentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#bucket_region TfSmbFileShare#bucket_region}.
	// Experimental.
	BucketRegion *string `field:"optional" json:"bucketRegion" yaml:"bucketRegion"`
	// cache_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#cache_attributes TfSmbFileShare#cache_attributes}
	// Experimental.
	CacheAttributes *TfSmbFileShare_CacheAttributesProperty `field:"optional" json:"cacheAttributes" yaml:"cacheAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#case_sensitivity TfSmbFileShare#case_sensitivity}.
	// Experimental.
	CaseSensitivity *string `field:"optional" json:"caseSensitivity" yaml:"caseSensitivity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#default_storage_class TfSmbFileShare#default_storage_class}.
	// Experimental.
	DefaultStorageClass *string `field:"optional" json:"defaultStorageClass" yaml:"defaultStorageClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#file_share_name TfSmbFileShare#file_share_name}.
	// Experimental.
	FileShareName *string `field:"optional" json:"fileShareName" yaml:"fileShareName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#guess_mime_type_enabled TfSmbFileShare#guess_mime_type_enabled}.
	// Experimental.
	GuessMimeTypeEnabled interface{} `field:"optional" json:"guessMimeTypeEnabled" yaml:"guessMimeTypeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#id TfSmbFileShare#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#invalid_user_list TfSmbFileShare#invalid_user_list}.
	// Experimental.
	InvalidUserList *[]*string `field:"optional" json:"invalidUserList" yaml:"invalidUserList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#kms_encrypted TfSmbFileShare#kms_encrypted}.
	// Experimental.
	KmsEncrypted interface{} `field:"optional" json:"kmsEncrypted" yaml:"kmsEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#kms_key_arn TfSmbFileShare#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#notification_policy TfSmbFileShare#notification_policy}.
	// Experimental.
	NotificationPolicy *string `field:"optional" json:"notificationPolicy" yaml:"notificationPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#object_acl TfSmbFileShare#object_acl}.
	// Experimental.
	ObjectAcl *string `field:"optional" json:"objectAcl" yaml:"objectAcl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#oplocks_enabled TfSmbFileShare#oplocks_enabled}.
	// Experimental.
	OplocksEnabled interface{} `field:"optional" json:"oplocksEnabled" yaml:"oplocksEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#read_only TfSmbFileShare#read_only}.
	// Experimental.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#region TfSmbFileShare#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#requester_pays TfSmbFileShare#requester_pays}.
	// Experimental.
	RequesterPays interface{} `field:"optional" json:"requesterPays" yaml:"requesterPays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#smb_acl_enabled TfSmbFileShare#smb_acl_enabled}.
	// Experimental.
	SmbAclEnabled interface{} `field:"optional" json:"smbAclEnabled" yaml:"smbAclEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#tags TfSmbFileShare#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#tags_all TfSmbFileShare#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#timeouts TfSmbFileShare#timeouts}
	// Experimental.
	Timeouts *TfSmbFileShare_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#valid_user_list TfSmbFileShare#valid_user_list}.
	// Experimental.
	ValidUserList *[]*string `field:"optional" json:"validUserList" yaml:"validUserList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#vpc_endpoint_dns_name TfSmbFileShare#vpc_endpoint_dns_name}.
	// Experimental.
	VpcEndpointDnsName *string `field:"optional" json:"vpcEndpointDnsName" yaml:"vpcEndpointDnsName"`
}

