package glue


// Experimental.
type AwsPartition_StorageDescriptorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#additional_locations AwsPartition#additional_locations}.
	// Experimental.
	AdditionalLocations *[]*string `field:"optional" json:"additionalLocations" yaml:"additionalLocations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#bucket_columns AwsPartition#bucket_columns}.
	// Experimental.
	BucketColumns *[]*string `field:"optional" json:"bucketColumns" yaml:"bucketColumns"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#columns AwsPartition#columns}
	// Experimental.
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#compressed AwsPartition#compressed}.
	// Experimental.
	Compressed interface{} `field:"optional" json:"compressed" yaml:"compressed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#input_format AwsPartition#input_format}.
	// Experimental.
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#location AwsPartition#location}.
	// Experimental.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#number_of_buckets AwsPartition#number_of_buckets}.
	// Experimental.
	NumberOfBuckets *float64 `field:"optional" json:"numberOfBuckets" yaml:"numberOfBuckets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#output_format AwsPartition#output_format}.
	// Experimental.
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#parameters AwsPartition#parameters}.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// ser_de_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#ser_de_info AwsPartition#ser_de_info}
	// Experimental.
	SerDeInfo *AwsPartition_SerDeInfoProperty `field:"optional" json:"serDeInfo" yaml:"serDeInfo"`
	// skewed_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#skewed_info AwsPartition#skewed_info}
	// Experimental.
	SkewedInfo *AwsPartition_SkewedInfoProperty `field:"optional" json:"skewedInfo" yaml:"skewedInfo"`
	// sort_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#sort_columns AwsPartition#sort_columns}
	// Experimental.
	SortColumns interface{} `field:"optional" json:"sortColumns" yaml:"sortColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#stored_as_sub_directories AwsPartition#stored_as_sub_directories}.
	// Experimental.
	StoredAsSubDirectories interface{} `field:"optional" json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
}

