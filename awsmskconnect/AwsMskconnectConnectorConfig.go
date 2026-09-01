package awsmskconnect

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMskconnectConnectorConfig struct {
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
	// capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#capacity AwsMskconnectConnector#capacity}
	// Experimental.
	Capacity *AwsMskconnectConnector_CapacityProperty `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#connector_configuration AwsMskconnectConnector#connector_configuration}.
	// Experimental.
	ConnectorConfiguration *map[string]*string `field:"required" json:"connectorConfiguration" yaml:"connectorConfiguration"`
	// kafka_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#kafka_cluster AwsMskconnectConnector#kafka_cluster}
	// Experimental.
	KafkaCluster *AwsMskconnectConnector_KafkaClusterProperty `field:"required" json:"kafkaCluster" yaml:"kafkaCluster"`
	// kafka_cluster_client_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#kafka_cluster_client_authentication AwsMskconnectConnector#kafka_cluster_client_authentication}
	// Experimental.
	KafkaClusterClientAuthentication *AwsMskconnectConnector_KafkaClusterClientAuthenticationProperty `field:"required" json:"kafkaClusterClientAuthentication" yaml:"kafkaClusterClientAuthentication"`
	// kafka_cluster_encryption_in_transit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#kafka_cluster_encryption_in_transit AwsMskconnectConnector#kafka_cluster_encryption_in_transit}
	// Experimental.
	KafkaClusterEncryptionInTransit *AwsMskconnectConnector_KafkaClusterEncryptionInTransitProperty `field:"required" json:"kafkaClusterEncryptionInTransit" yaml:"kafkaClusterEncryptionInTransit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#kafkaconnect_version AwsMskconnectConnector#kafkaconnect_version}.
	// Experimental.
	KafkaconnectVersion *string `field:"required" json:"kafkaconnectVersion" yaml:"kafkaconnectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#name AwsMskconnectConnector#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// plugin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#plugin AwsMskconnectConnector#plugin}
	// Experimental.
	Plugin interface{} `field:"required" json:"plugin" yaml:"plugin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#service_execution_role_arn AwsMskconnectConnector#service_execution_role_arn}.
	// Experimental.
	ServiceExecutionRoleArn *string `field:"required" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#description AwsMskconnectConnector#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#id AwsMskconnectConnector#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// log_delivery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#log_delivery AwsMskconnectConnector#log_delivery}
	// Experimental.
	LogDelivery *AwsMskconnectConnector_LogDeliveryProperty `field:"optional" json:"logDelivery" yaml:"logDelivery"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#region AwsMskconnectConnector#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#tags AwsMskconnectConnector#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#tags_all AwsMskconnectConnector#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#timeouts AwsMskconnectConnector#timeouts}
	// Experimental.
	Timeouts *AwsMskconnectConnector_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// worker_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#worker_configuration AwsMskconnectConnector#worker_configuration}
	// Experimental.
	WorkerConfiguration *AwsMskconnectConnector_WorkerConfigurationProperty `field:"optional" json:"workerConfiguration" yaml:"workerConfiguration"`
}

