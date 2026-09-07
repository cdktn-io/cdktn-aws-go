package bedrockagents


// Experimental.
type AwsKnowledgeBase_StorageConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#type AwsKnowledgeBase#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// mongo_db_atlas_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#mongo_db_atlas_configuration AwsKnowledgeBase#mongo_db_atlas_configuration}
	// Experimental.
	MongoDbAtlasConfiguration interface{} `field:"optional" json:"mongoDbAtlasConfiguration" yaml:"mongoDbAtlasConfiguration"`
	// neptune_analytics_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#neptune_analytics_configuration AwsKnowledgeBase#neptune_analytics_configuration}
	// Experimental.
	NeptuneAnalyticsConfiguration interface{} `field:"optional" json:"neptuneAnalyticsConfiguration" yaml:"neptuneAnalyticsConfiguration"`
	// opensearch_managed_cluster_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#opensearch_managed_cluster_configuration AwsKnowledgeBase#opensearch_managed_cluster_configuration}
	// Experimental.
	OpensearchManagedClusterConfiguration interface{} `field:"optional" json:"opensearchManagedClusterConfiguration" yaml:"opensearchManagedClusterConfiguration"`
	// opensearch_serverless_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#opensearch_serverless_configuration AwsKnowledgeBase#opensearch_serverless_configuration}
	// Experimental.
	OpensearchServerlessConfiguration interface{} `field:"optional" json:"opensearchServerlessConfiguration" yaml:"opensearchServerlessConfiguration"`
	// pinecone_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#pinecone_configuration AwsKnowledgeBase#pinecone_configuration}
	// Experimental.
	PineconeConfiguration interface{} `field:"optional" json:"pineconeConfiguration" yaml:"pineconeConfiguration"`
	// rds_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#rds_configuration AwsKnowledgeBase#rds_configuration}
	// Experimental.
	RdsConfiguration interface{} `field:"optional" json:"rdsConfiguration" yaml:"rdsConfiguration"`
	// redis_enterprise_cloud_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#redis_enterprise_cloud_configuration AwsKnowledgeBase#redis_enterprise_cloud_configuration}
	// Experimental.
	RedisEnterpriseCloudConfiguration interface{} `field:"optional" json:"redisEnterpriseCloudConfiguration" yaml:"redisEnterpriseCloudConfiguration"`
	// s3_vectors_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#s3_vectors_configuration AwsKnowledgeBase#s3_vectors_configuration}
	// Experimental.
	S3VectorsConfiguration interface{} `field:"optional" json:"s3VectorsConfiguration" yaml:"s3VectorsConfiguration"`
}

