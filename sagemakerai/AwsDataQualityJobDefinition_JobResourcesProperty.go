package sagemakerai


// Experimental.
type AwsDataQualityJobDefinition_JobResourcesProperty struct {
	// cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#cluster_config AwsDataQualityJobDefinition#cluster_config}
	// Experimental.
	ClusterConfig *AwsDataQualityJobDefinition_ClusterConfigProperty `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

