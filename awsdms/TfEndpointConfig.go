package awsdms

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#endpoint_id TfEndpoint#endpoint_id}.
	// Experimental.
	EndpointId *string `field:"required" json:"endpointId" yaml:"endpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#endpoint_type TfEndpoint#endpoint_type}.
	// Experimental.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#engine_name TfEndpoint#engine_name}.
	// Experimental.
	EngineName *string `field:"required" json:"engineName" yaml:"engineName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#certificate_arn TfEndpoint#certificate_arn}.
	// Experimental.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#database_name TfEndpoint#database_name}.
	// Experimental.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// elasticsearch_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#elasticsearch_settings TfEndpoint#elasticsearch_settings}
	// Experimental.
	ElasticsearchSettings *TfEndpoint_ElasticsearchSettingsProperty `field:"optional" json:"elasticsearchSettings" yaml:"elasticsearchSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#extra_connection_attributes TfEndpoint#extra_connection_attributes}.
	// Experimental.
	ExtraConnectionAttributes *string `field:"optional" json:"extraConnectionAttributes" yaml:"extraConnectionAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#id TfEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kafka_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#kafka_settings TfEndpoint#kafka_settings}
	// Experimental.
	KafkaSettings *TfEndpoint_KafkaSettingsProperty `field:"optional" json:"kafkaSettings" yaml:"kafkaSettings"`
	// kinesis_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#kinesis_settings TfEndpoint#kinesis_settings}
	// Experimental.
	KinesisSettings *TfEndpoint_KinesisSettingsProperty `field:"optional" json:"kinesisSettings" yaml:"kinesisSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#kms_key_arn TfEndpoint#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// mongodb_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#mongodb_settings TfEndpoint#mongodb_settings}
	// Experimental.
	MongodbSettings *TfEndpoint_MongodbSettingsProperty `field:"optional" json:"mongodbSettings" yaml:"mongodbSettings"`
	// mysql_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#mysql_settings TfEndpoint#mysql_settings}
	// Experimental.
	MysqlSettings *TfEndpoint_MysqlSettingsProperty `field:"optional" json:"mysqlSettings" yaml:"mysqlSettings"`
	// oracle_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#oracle_settings TfEndpoint#oracle_settings}
	// Experimental.
	OracleSettings *TfEndpoint_OracleSettingsProperty `field:"optional" json:"oracleSettings" yaml:"oracleSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#password TfEndpoint#password}.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#pause_replication_tasks TfEndpoint#pause_replication_tasks}.
	// Experimental.
	PauseReplicationTasks interface{} `field:"optional" json:"pauseReplicationTasks" yaml:"pauseReplicationTasks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#port TfEndpoint#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// postgres_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#postgres_settings TfEndpoint#postgres_settings}
	// Experimental.
	PostgresSettings *TfEndpoint_PostgresSettingsProperty `field:"optional" json:"postgresSettings" yaml:"postgresSettings"`
	// redis_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#redis_settings TfEndpoint#redis_settings}
	// Experimental.
	RedisSettings *TfEndpoint_RedisSettingsProperty `field:"optional" json:"redisSettings" yaml:"redisSettings"`
	// redshift_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#redshift_settings TfEndpoint#redshift_settings}
	// Experimental.
	RedshiftSettings *TfEndpoint_RedshiftSettingsProperty `field:"optional" json:"redshiftSettings" yaml:"redshiftSettings"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#region TfEndpoint#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#secrets_manager_access_role_arn TfEndpoint#secrets_manager_access_role_arn}.
	// Experimental.
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#secrets_manager_arn TfEndpoint#secrets_manager_arn}.
	// Experimental.
	SecretsManagerArn *string `field:"optional" json:"secretsManagerArn" yaml:"secretsManagerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#server_name TfEndpoint#server_name}.
	// Experimental.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#service_access_role TfEndpoint#service_access_role}.
	// Experimental.
	ServiceAccessRole *string `field:"optional" json:"serviceAccessRole" yaml:"serviceAccessRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#ssl_mode TfEndpoint#ssl_mode}.
	// Experimental.
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#tags TfEndpoint#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#tags_all TfEndpoint#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#timeouts TfEndpoint#timeouts}
	// Experimental.
	Timeouts *TfEndpoint_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#username TfEndpoint#username}.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

