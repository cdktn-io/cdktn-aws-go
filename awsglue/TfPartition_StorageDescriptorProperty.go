package awsglue


// Experimental.
type TfPartition_StorageDescriptorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#additional_locations TfPartition#additional_locations}.
	// Experimental.
	AdditionalLocations *[]*string `field:"optional" json:"additionalLocations" yaml:"additionalLocations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#bucket_columns TfPartition#bucket_columns}.
	// Experimental.
	BucketColumns *[]*string `field:"optional" json:"bucketColumns" yaml:"bucketColumns"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#columns TfPartition#columns}
	// Experimental.
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#compressed TfPartition#compressed}.
	// Experimental.
	Compressed interface{} `field:"optional" json:"compressed" yaml:"compressed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#input_format TfPartition#input_format}.
	// Experimental.
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#location TfPartition#location}.
	// Experimental.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#number_of_buckets TfPartition#number_of_buckets}.
	// Experimental.
	NumberOfBuckets *float64 `field:"optional" json:"numberOfBuckets" yaml:"numberOfBuckets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#output_format TfPartition#output_format}.
	// Experimental.
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#parameters TfPartition#parameters}.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// ser_de_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#ser_de_info TfPartition#ser_de_info}
	// Experimental.
	SerDeInfo *TfPartition_SerDeInfoProperty `field:"optional" json:"serDeInfo" yaml:"serDeInfo"`
	// skewed_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#skewed_info TfPartition#skewed_info}
	// Experimental.
	SkewedInfo *TfPartition_SkewedInfoProperty `field:"optional" json:"skewedInfo" yaml:"skewedInfo"`
	// sort_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#sort_columns TfPartition#sort_columns}
	// Experimental.
	SortColumns interface{} `field:"optional" json:"sortColumns" yaml:"sortColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#stored_as_sub_directories TfPartition#stored_as_sub_directories}.
	// Experimental.
	StoredAsSubDirectories interface{} `field:"optional" json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
}

