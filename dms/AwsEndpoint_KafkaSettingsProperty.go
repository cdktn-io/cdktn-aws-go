package dms


// Experimental.
type AwsEndpoint_KafkaSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#broker AwsEndpoint#broker}.
	// Experimental.
	Broker *string `field:"required" json:"broker" yaml:"broker"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#include_control_details AwsEndpoint#include_control_details}.
	// Experimental.
	IncludeControlDetails interface{} `field:"optional" json:"includeControlDetails" yaml:"includeControlDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#include_null_and_empty AwsEndpoint#include_null_and_empty}.
	// Experimental.
	IncludeNullAndEmpty interface{} `field:"optional" json:"includeNullAndEmpty" yaml:"includeNullAndEmpty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#include_partition_value AwsEndpoint#include_partition_value}.
	// Experimental.
	IncludePartitionValue interface{} `field:"optional" json:"includePartitionValue" yaml:"includePartitionValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#include_table_alter_operations AwsEndpoint#include_table_alter_operations}.
	// Experimental.
	IncludeTableAlterOperations interface{} `field:"optional" json:"includeTableAlterOperations" yaml:"includeTableAlterOperations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#include_transaction_details AwsEndpoint#include_transaction_details}.
	// Experimental.
	IncludeTransactionDetails interface{} `field:"optional" json:"includeTransactionDetails" yaml:"includeTransactionDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#message_format AwsEndpoint#message_format}.
	// Experimental.
	MessageFormat *string `field:"optional" json:"messageFormat" yaml:"messageFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#message_max_bytes AwsEndpoint#message_max_bytes}.
	// Experimental.
	MessageMaxBytes *float64 `field:"optional" json:"messageMaxBytes" yaml:"messageMaxBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#no_hex_prefix AwsEndpoint#no_hex_prefix}.
	// Experimental.
	NoHexPrefix interface{} `field:"optional" json:"noHexPrefix" yaml:"noHexPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#partition_include_schema_table AwsEndpoint#partition_include_schema_table}.
	// Experimental.
	PartitionIncludeSchemaTable interface{} `field:"optional" json:"partitionIncludeSchemaTable" yaml:"partitionIncludeSchemaTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#sasl_mechanism AwsEndpoint#sasl_mechanism}.
	// Experimental.
	SaslMechanism *string `field:"optional" json:"saslMechanism" yaml:"saslMechanism"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#sasl_password AwsEndpoint#sasl_password}.
	// Experimental.
	SaslPassword *string `field:"optional" json:"saslPassword" yaml:"saslPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#sasl_username AwsEndpoint#sasl_username}.
	// Experimental.
	SaslUsername *string `field:"optional" json:"saslUsername" yaml:"saslUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#security_protocol AwsEndpoint#security_protocol}.
	// Experimental.
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#ssl_ca_certificate_arn AwsEndpoint#ssl_ca_certificate_arn}.
	// Experimental.
	SslCaCertificateArn *string `field:"optional" json:"sslCaCertificateArn" yaml:"sslCaCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#ssl_client_certificate_arn AwsEndpoint#ssl_client_certificate_arn}.
	// Experimental.
	SslClientCertificateArn *string `field:"optional" json:"sslClientCertificateArn" yaml:"sslClientCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#ssl_client_key_arn AwsEndpoint#ssl_client_key_arn}.
	// Experimental.
	SslClientKeyArn *string `field:"optional" json:"sslClientKeyArn" yaml:"sslClientKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#ssl_client_key_password AwsEndpoint#ssl_client_key_password}.
	// Experimental.
	SslClientKeyPassword *string `field:"optional" json:"sslClientKeyPassword" yaml:"sslClientKeyPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#topic AwsEndpoint#topic}.
	// Experimental.
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

