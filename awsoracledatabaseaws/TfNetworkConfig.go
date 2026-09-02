package awsoracledatabaseaws

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfNetworkConfig struct {
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
	// The AZ ID of the AZ where the ODB network is located.
	//
	// Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#availability_zone_id TfNetwork#availability_zone_id}
	// Experimental.
	AvailabilityZoneId *string `field:"required" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// The CIDR range of the backup subnet for the ODB network.
	//
	// Changing this will force terraform to create new resource.
	// 	Constraints:
	// 	   - Must not overlap with the CIDR range of the client subnet.
	// 	   - Must not overlap with the CIDR ranges of the VPCs that are connected to the
	// 	   ODB network.
	// 	   - Must not use the following CIDR ranges that are reserved by OCI:
	// 	   - 100.106.0.0/16 and 100.107.0.0/16
	// 	   - 169.254.0.0/16
	// 	   - 224.0.0.0 - 239.255.255.255
	// 	   - 240.0.0.0 - 255.255.255.255
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#backup_subnet_cidr TfNetwork#backup_subnet_cidr}
	// Experimental.
	BackupSubnetCidr *string `field:"required" json:"backupSubnetCidr" yaml:"backupSubnetCidr"`
	// The CIDR notation for the network resource.
	//
	// Changing this will force terraform to create new resource.
	//  Constraints:
	//   	 - Must not overlap with the CIDR range of the backup subnet.
	//    	- Must not overlap with the CIDR ranges of the VPCs that are connected to the
	//    ODB network.
	//   	- Must not use the following CIDR ranges that are reserved by OCI:
	//   	 - 100.106.0.0/16 and 100.107.0.0/16
	//   	 - 169.254.0.0/16
	//    	- 224.0.0.0 - 239.255.255.255
	//    	- 240.0.0.0 - 255.255.255.255
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#client_subnet_cidr TfNetwork#client_subnet_cidr}
	// Experimental.
	ClientSubnetCidr *string `field:"required" json:"clientSubnetCidr" yaml:"clientSubnetCidr"`
	// The user-friendly name for the odb network. Changing this will force terraform to create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#display_name TfNetwork#display_name}
	// Experimental.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Specifies the configuration for Amazon S3 access from the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#s3_access TfNetwork#s3_access}
	// Experimental.
	S3Access *string `field:"required" json:"s3Access" yaml:"s3Access"`
	// Specifies the configuration for Zero-ETL access from the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#zero_etl_access TfNetwork#zero_etl_access}
	// Experimental.
	ZeroEtlAccess *string `field:"required" json:"zeroEtlAccess" yaml:"zeroEtlAccess"`
	// The name of the Availability Zone (AZ) where the odb network is located.
	//
	// Changing this will force terraform to create new resource
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#availability_zone TfNetwork#availability_zone}
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The list of regions enabled for cross-region restore in the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#cross_region_s3_restore_sources_access TfNetwork#cross_region_s3_restore_sources_access}
	// Experimental.
	CrossRegionS3RestoreSourcesAccess *[]*string `field:"optional" json:"crossRegionS3RestoreSourcesAccess" yaml:"crossRegionS3RestoreSourcesAccess"`
	// The name of the custom domain that the network is located. custom_domain_name and default_dns_prefix both can't be given.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#custom_domain_name TfNetwork#custom_domain_name}
	// Experimental.
	CustomDomainName *string `field:"optional" json:"customDomainName" yaml:"customDomainName"`
	// The default DNS prefix for the network resource. Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#default_dns_prefix TfNetwork#default_dns_prefix}
	// Experimental.
	DefaultDnsPrefix *string `field:"optional" json:"defaultDnsPrefix" yaml:"defaultDnsPrefix"`
	// If set to true deletes associated OCI resources. Default false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#delete_associated_resources TfNetwork#delete_associated_resources}
	// Experimental.
	DeleteAssociatedResources interface{} `field:"optional" json:"deleteAssociatedResources" yaml:"deleteAssociatedResources"`
	// Specifies the configuration for Amazon KMS access from the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#kms_access TfNetwork#kms_access}
	// Experimental.
	KmsAccess *string `field:"optional" json:"kmsAccess" yaml:"kmsAccess"`
	// Specifies the endpoint policy for Amazon KMS access from the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#kms_policy_document TfNetwork#kms_policy_document}
	// Experimental.
	KmsPolicyDocument *string `field:"optional" json:"kmsPolicyDocument" yaml:"kmsPolicyDocument"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#region TfNetwork#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Specifies the endpoint policy for Amazon S3 access from the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#s3_policy_document TfNetwork#s3_policy_document}
	// Experimental.
	S3PolicyDocument *string `field:"optional" json:"s3PolicyDocument" yaml:"s3PolicyDocument"`
	// Specifies the configuration for Amazon STS access from the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#sts_access TfNetwork#sts_access}
	// Experimental.
	StsAccess *string `field:"optional" json:"stsAccess" yaml:"stsAccess"`
	// Specifies the endpoint policy for Amazon STS access from the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#sts_policy_document TfNetwork#sts_policy_document}
	// Experimental.
	StsPolicyDocument *string `field:"optional" json:"stsPolicyDocument" yaml:"stsPolicyDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#tags TfNetwork#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network#timeouts TfNetwork#timeouts}
	// Experimental.
	Timeouts *TfNetwork_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

