package emr


// Experimental.
type AwsCluster_StepProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#action_on_failure AwsCluster#action_on_failure}.
	// Experimental.
	ActionOnFailure *string `field:"optional" json:"actionOnFailure" yaml:"actionOnFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#hadoop_jar_step AwsCluster#hadoop_jar_step}.
	// Experimental.
	HadoopJarStep interface{} `field:"optional" json:"hadoopJarStep" yaml:"hadoopJarStep"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#name AwsCluster#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

