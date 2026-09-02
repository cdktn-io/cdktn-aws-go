package awsmq

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfBrokerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#broker_name TfBroker#broker_name}.
	// Experimental.
	BrokerName *string `field:"required" json:"brokerName" yaml:"brokerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#engine_type TfBroker#engine_type}.
	// Experimental.
	EngineType *string `field:"required" json:"engineType" yaml:"engineType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#engine_version TfBroker#engine_version}.
	// Experimental.
	EngineVersion *string `field:"required" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#host_instance_type TfBroker#host_instance_type}.
	// Experimental.
	HostInstanceType *string `field:"required" json:"hostInstanceType" yaml:"hostInstanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#apply_immediately TfBroker#apply_immediately}.
	// Experimental.
	ApplyImmediately interface{} `field:"optional" json:"applyImmediately" yaml:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#authentication_strategy TfBroker#authentication_strategy}.
	// Experimental.
	AuthenticationStrategy *string `field:"optional" json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#auto_minor_version_upgrade TfBroker#auto_minor_version_upgrade}.
	// Experimental.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#configuration TfBroker#configuration}
	// Experimental.
	Configuration *TfBroker_ConfigurationProperty `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#data_replication_mode TfBroker#data_replication_mode}.
	// Experimental.
	DataReplicationMode *string `field:"optional" json:"dataReplicationMode" yaml:"dataReplicationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#data_replication_primary_broker_arn TfBroker#data_replication_primary_broker_arn}.
	// Experimental.
	DataReplicationPrimaryBrokerArn *string `field:"optional" json:"dataReplicationPrimaryBrokerArn" yaml:"dataReplicationPrimaryBrokerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#deployment_mode TfBroker#deployment_mode}.
	// Experimental.
	DeploymentMode *string `field:"optional" json:"deploymentMode" yaml:"deploymentMode"`
	// encryption_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#encryption_options TfBroker#encryption_options}
	// Experimental.
	EncryptionOptions *TfBroker_EncryptionOptionsProperty `field:"optional" json:"encryptionOptions" yaml:"encryptionOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#id TfBroker#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ldap_server_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#ldap_server_metadata TfBroker#ldap_server_metadata}
	// Experimental.
	LdapServerMetadata *TfBroker_LdapServerMetadataProperty `field:"optional" json:"ldapServerMetadata" yaml:"ldapServerMetadata"`
	// logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#logs TfBroker#logs}
	// Experimental.
	Logs *TfBroker_LogsProperty `field:"optional" json:"logs" yaml:"logs"`
	// maintenance_window_start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#maintenance_window_start_time TfBroker#maintenance_window_start_time}
	// Experimental.
	MaintenanceWindowStartTime *TfBroker_MaintenanceWindowStartTimeProperty `field:"optional" json:"maintenanceWindowStartTime" yaml:"maintenanceWindowStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#publicly_accessible TfBroker#publicly_accessible}.
	// Experimental.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#region TfBroker#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#resource_share_arns TfBroker#resource_share_arns}.
	// Experimental.
	ResourceShareArns *[]*string `field:"optional" json:"resourceShareArns" yaml:"resourceShareArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#security_groups TfBroker#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#storage_type TfBroker#storage_type}.
	// Experimental.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#subnet_ids TfBroker#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#tags TfBroker#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#tags_all TfBroker#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#timeouts TfBroker#timeouts}
	// Experimental.
	Timeouts *TfBroker_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#user TfBroker#user}
	// Experimental.
	User interface{} `field:"optional" json:"user" yaml:"user"`
}

