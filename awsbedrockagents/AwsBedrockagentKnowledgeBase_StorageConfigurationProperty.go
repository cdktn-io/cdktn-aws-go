package awsbedrockagents


// Experimental.
type AwsBedrockagentKnowledgeBase_StorageConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#type AwsBedrockagentKnowledgeBase#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// mongo_db_atlas_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#mongo_db_atlas_configuration AwsBedrockagentKnowledgeBase#mongo_db_atlas_configuration}
	// Experimental.
	MongoDbAtlasConfiguration interface{} `field:"optional" json:"mongoDbAtlasConfiguration" yaml:"mongoDbAtlasConfiguration"`
	// neptune_analytics_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#neptune_analytics_configuration AwsBedrockagentKnowledgeBase#neptune_analytics_configuration}
	// Experimental.
	NeptuneAnalyticsConfiguration interface{} `field:"optional" json:"neptuneAnalyticsConfiguration" yaml:"neptuneAnalyticsConfiguration"`
	// opensearch_managed_cluster_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#opensearch_managed_cluster_configuration AwsBedrockagentKnowledgeBase#opensearch_managed_cluster_configuration}
	// Experimental.
	OpensearchManagedClusterConfiguration interface{} `field:"optional" json:"opensearchManagedClusterConfiguration" yaml:"opensearchManagedClusterConfiguration"`
	// opensearch_serverless_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#opensearch_serverless_configuration AwsBedrockagentKnowledgeBase#opensearch_serverless_configuration}
	// Experimental.
	OpensearchServerlessConfiguration interface{} `field:"optional" json:"opensearchServerlessConfiguration" yaml:"opensearchServerlessConfiguration"`
	// pinecone_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#pinecone_configuration AwsBedrockagentKnowledgeBase#pinecone_configuration}
	// Experimental.
	PineconeConfiguration interface{} `field:"optional" json:"pineconeConfiguration" yaml:"pineconeConfiguration"`
	// rds_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#rds_configuration AwsBedrockagentKnowledgeBase#rds_configuration}
	// Experimental.
	RdsConfiguration interface{} `field:"optional" json:"rdsConfiguration" yaml:"rdsConfiguration"`
	// redis_enterprise_cloud_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#redis_enterprise_cloud_configuration AwsBedrockagentKnowledgeBase#redis_enterprise_cloud_configuration}
	// Experimental.
	RedisEnterpriseCloudConfiguration interface{} `field:"optional" json:"redisEnterpriseCloudConfiguration" yaml:"redisEnterpriseCloudConfiguration"`
	// s3_vectors_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#s3_vectors_configuration AwsBedrockagentKnowledgeBase#s3_vectors_configuration}
	// Experimental.
	S3VectorsConfiguration interface{} `field:"optional" json:"s3VectorsConfiguration" yaml:"s3VectorsConfiguration"`
}

