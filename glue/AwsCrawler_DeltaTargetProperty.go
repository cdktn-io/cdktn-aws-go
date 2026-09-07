package glue


// Experimental.
type AwsCrawler_DeltaTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#delta_tables AwsCrawler#delta_tables}.
	// Experimental.
	DeltaTables *[]*string `field:"required" json:"deltaTables" yaml:"deltaTables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#write_manifest AwsCrawler#write_manifest}.
	// Experimental.
	WriteManifest interface{} `field:"required" json:"writeManifest" yaml:"writeManifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#connection_name AwsCrawler#connection_name}.
	// Experimental.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#create_native_delta_table AwsCrawler#create_native_delta_table}.
	// Experimental.
	CreateNativeDeltaTable interface{} `field:"optional" json:"createNativeDeltaTable" yaml:"createNativeDeltaTable"`
}

