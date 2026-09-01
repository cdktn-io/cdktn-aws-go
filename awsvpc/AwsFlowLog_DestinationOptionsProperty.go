package awsvpc


// Experimental.
type AwsFlowLog_DestinationOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#file_format AwsFlowLog#file_format}.
	// Experimental.
	FileFormat *string `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#hive_compatible_partitions AwsFlowLog#hive_compatible_partitions}.
	// Experimental.
	HiveCompatiblePartitions interface{} `field:"optional" json:"hiveCompatiblePartitions" yaml:"hiveCompatiblePartitions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#per_hour_partition AwsFlowLog#per_hour_partition}.
	// Experimental.
	PerHourPartition interface{} `field:"optional" json:"perHourPartition" yaml:"perHourPartition"`
}

