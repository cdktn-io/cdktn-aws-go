package bcmdataexports


// Experimental.
type AwsExport_DataQueryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#query_statement AwsExport#query_statement}.
	// Experimental.
	QueryStatement *string `field:"required" json:"queryStatement" yaml:"queryStatement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#table_configurations AwsExport#table_configurations}.
	// Experimental.
	TableConfigurations interface{} `field:"optional" json:"tableConfigurations" yaml:"tableConfigurations"`
}

