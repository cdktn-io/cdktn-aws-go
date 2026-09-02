package awsdms

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReplicationInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#replication_instance_class TfReplicationInstance#replication_instance_class}.
	// Experimental.
	ReplicationInstanceClass *string `field:"required" json:"replicationInstanceClass" yaml:"replicationInstanceClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#replication_instance_id TfReplicationInstance#replication_instance_id}.
	// Experimental.
	ReplicationInstanceId *string `field:"required" json:"replicationInstanceId" yaml:"replicationInstanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#allocated_storage TfReplicationInstance#allocated_storage}.
	// Experimental.
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#allow_major_version_upgrade TfReplicationInstance#allow_major_version_upgrade}.
	// Experimental.
	AllowMajorVersionUpgrade interface{} `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#apply_immediately TfReplicationInstance#apply_immediately}.
	// Experimental.
	ApplyImmediately interface{} `field:"optional" json:"applyImmediately" yaml:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#auto_minor_version_upgrade TfReplicationInstance#auto_minor_version_upgrade}.
	// Experimental.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#availability_zone TfReplicationInstance#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#dns_name_servers TfReplicationInstance#dns_name_servers}.
	// Experimental.
	DnsNameServers *string `field:"optional" json:"dnsNameServers" yaml:"dnsNameServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#engine_version TfReplicationInstance#engine_version}.
	// Experimental.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#id TfReplicationInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kerberos_authentication_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#kerberos_authentication_settings TfReplicationInstance#kerberos_authentication_settings}
	// Experimental.
	KerberosAuthenticationSettings *TfReplicationInstance_KerberosAuthenticationSettingsProperty `field:"optional" json:"kerberosAuthenticationSettings" yaml:"kerberosAuthenticationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#kms_key_arn TfReplicationInstance#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#multi_az TfReplicationInstance#multi_az}.
	// Experimental.
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#network_type TfReplicationInstance#network_type}.
	// Experimental.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#preferred_maintenance_window TfReplicationInstance#preferred_maintenance_window}.
	// Experimental.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#publicly_accessible TfReplicationInstance#publicly_accessible}.
	// Experimental.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#region TfReplicationInstance#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#replication_subnet_group_id TfReplicationInstance#replication_subnet_group_id}.
	// Experimental.
	ReplicationSubnetGroupId *string `field:"optional" json:"replicationSubnetGroupId" yaml:"replicationSubnetGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#tags TfReplicationInstance#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#tags_all TfReplicationInstance#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#timeouts TfReplicationInstance#timeouts}
	// Experimental.
	Timeouts *TfReplicationInstance_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#vpc_security_group_ids TfReplicationInstance#vpc_security_group_ids}.
	// Experimental.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

