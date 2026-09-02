package awselasticache


// Experimental.
type TfReplicationGroup_LogDeliveryConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#destination TfReplicationGroup#destination}.
	// Experimental.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#destination_type TfReplicationGroup#destination_type}.
	// Experimental.
	DestinationType *string `field:"required" json:"destinationType" yaml:"destinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#log_format TfReplicationGroup#log_format}.
	// Experimental.
	LogFormat *string `field:"required" json:"logFormat" yaml:"logFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#log_type TfReplicationGroup#log_type}.
	// Experimental.
	LogType *string `field:"required" json:"logType" yaml:"logType"`
}

