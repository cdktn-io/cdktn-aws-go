package awsglue


// Experimental.
type TfCatalogTable_StorageDescriptorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#additional_locations TfCatalogTable#additional_locations}.
	// Experimental.
	AdditionalLocations *[]*string `field:"optional" json:"additionalLocations" yaml:"additionalLocations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#bucket_columns TfCatalogTable#bucket_columns}.
	// Experimental.
	BucketColumns *[]*string `field:"optional" json:"bucketColumns" yaml:"bucketColumns"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#columns TfCatalogTable#columns}
	// Experimental.
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#compressed TfCatalogTable#compressed}.
	// Experimental.
	Compressed interface{} `field:"optional" json:"compressed" yaml:"compressed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#input_format TfCatalogTable#input_format}.
	// Experimental.
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#location TfCatalogTable#location}.
	// Experimental.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#number_of_buckets TfCatalogTable#number_of_buckets}.
	// Experimental.
	NumberOfBuckets *float64 `field:"optional" json:"numberOfBuckets" yaml:"numberOfBuckets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#output_format TfCatalogTable#output_format}.
	// Experimental.
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#parameters TfCatalogTable#parameters}.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// schema_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#schema_reference TfCatalogTable#schema_reference}
	// Experimental.
	SchemaReference *TfCatalogTable_SchemaReferenceProperty `field:"optional" json:"schemaReference" yaml:"schemaReference"`
	// ser_de_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#ser_de_info TfCatalogTable#ser_de_info}
	// Experimental.
	SerDeInfo *TfCatalogTable_SerDeInfoProperty `field:"optional" json:"serDeInfo" yaml:"serDeInfo"`
	// skewed_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#skewed_info TfCatalogTable#skewed_info}
	// Experimental.
	SkewedInfo *TfCatalogTable_SkewedInfoProperty `field:"optional" json:"skewedInfo" yaml:"skewedInfo"`
	// sort_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#sort_columns TfCatalogTable#sort_columns}
	// Experimental.
	SortColumns interface{} `field:"optional" json:"sortColumns" yaml:"sortColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#stored_as_sub_directories TfCatalogTable#stored_as_sub_directories}.
	// Experimental.
	StoredAsSubDirectories interface{} `field:"optional" json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
}

