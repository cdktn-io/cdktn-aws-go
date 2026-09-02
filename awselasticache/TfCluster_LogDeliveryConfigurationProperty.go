package awselasticache


// Experimental.
type TfCluster_LogDeliveryConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_cluster#destination TfCluster#destination}.
	// Experimental.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_cluster#destination_type TfCluster#destination_type}.
	// Experimental.
	DestinationType *string `field:"required" json:"destinationType" yaml:"destinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_cluster#log_format TfCluster#log_format}.
	// Experimental.
	LogFormat *string `field:"required" json:"logFormat" yaml:"logFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_cluster#log_type TfCluster#log_type}.
	// Experimental.
	LogType *string `field:"required" json:"logType" yaml:"logType"`
}

