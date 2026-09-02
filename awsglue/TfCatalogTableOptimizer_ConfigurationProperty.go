package awsglue


// Experimental.
type TfCatalogTableOptimizer_ConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#enabled TfCatalogTableOptimizer#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#role_arn TfCatalogTableOptimizer#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// compaction_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#compaction_configuration TfCatalogTableOptimizer#compaction_configuration}
	// Experimental.
	CompactionConfiguration interface{} `field:"optional" json:"compactionConfiguration" yaml:"compactionConfiguration"`
	// orphan_file_deletion_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#orphan_file_deletion_configuration TfCatalogTableOptimizer#orphan_file_deletion_configuration}
	// Experimental.
	OrphanFileDeletionConfiguration interface{} `field:"optional" json:"orphanFileDeletionConfiguration" yaml:"orphanFileDeletionConfiguration"`
	// retention_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#retention_configuration TfCatalogTableOptimizer#retention_configuration}
	// Experimental.
	RetentionConfiguration interface{} `field:"optional" json:"retentionConfiguration" yaml:"retentionConfiguration"`
}

