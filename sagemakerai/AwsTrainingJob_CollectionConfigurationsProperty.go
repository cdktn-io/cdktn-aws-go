package sagemakerai


// Experimental.
type AwsTrainingJob_CollectionConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#collection_name AwsTrainingJob#collection_name}.
	// Experimental.
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#collection_parameters AwsTrainingJob#collection_parameters}.
	// Experimental.
	CollectionParameters *map[string]*string `field:"optional" json:"collectionParameters" yaml:"collectionParameters"`
}

