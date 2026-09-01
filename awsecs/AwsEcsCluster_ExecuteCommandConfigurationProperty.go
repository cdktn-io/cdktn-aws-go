package awsecs


// Experimental.
type AwsEcsCluster_ExecuteCommandConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#kms_key_id AwsEcsCluster#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#log_configuration AwsEcsCluster#log_configuration}
	// Experimental.
	LogConfiguration *AwsEcsCluster_LogConfigurationProperty `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#logging AwsEcsCluster#logging}.
	// Experimental.
	Logging *string `field:"optional" json:"logging" yaml:"logging"`
}

