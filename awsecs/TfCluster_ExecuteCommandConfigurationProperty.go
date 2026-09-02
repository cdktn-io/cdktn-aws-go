package awsecs


// Experimental.
type TfCluster_ExecuteCommandConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#kms_key_id TfCluster#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#log_configuration TfCluster#log_configuration}
	// Experimental.
	LogConfiguration *TfCluster_LogConfigurationProperty `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#logging TfCluster#logging}.
	// Experimental.
	Logging *string `field:"optional" json:"logging" yaml:"logging"`
}

