package glue


// Experimental.
type AwsCatalog_IcebergOptimizationPropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#compaction AwsCatalog#compaction}.
	// Experimental.
	Compaction *map[string]*string `field:"optional" json:"compaction" yaml:"compaction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#orphan_file_deletion AwsCatalog#orphan_file_deletion}.
	// Experimental.
	OrphanFileDeletion *map[string]*string `field:"optional" json:"orphanFileDeletion" yaml:"orphanFileDeletion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#retention AwsCatalog#retention}.
	// Experimental.
	Retention *map[string]*string `field:"optional" json:"retention" yaml:"retention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#role_arn AwsCatalog#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

