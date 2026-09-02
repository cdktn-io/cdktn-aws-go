package awsglue


// Experimental.
type TfCatalog_IcebergOptimizationPropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#compaction TfCatalog#compaction}.
	// Experimental.
	Compaction *map[string]*string `field:"optional" json:"compaction" yaml:"compaction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#orphan_file_deletion TfCatalog#orphan_file_deletion}.
	// Experimental.
	OrphanFileDeletion *map[string]*string `field:"optional" json:"orphanFileDeletion" yaml:"orphanFileDeletion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#retention TfCatalog#retention}.
	// Experimental.
	Retention *map[string]*string `field:"optional" json:"retention" yaml:"retention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#role_arn TfCatalog#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

