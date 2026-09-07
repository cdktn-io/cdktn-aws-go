package dynamodb


// Experimental.
type AwsTable_GlobalSecondaryIndexProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#name AwsTable#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#projection_type AwsTable#projection_type}.
	// Experimental.
	ProjectionType *string `field:"required" json:"projectionType" yaml:"projectionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#hash_key AwsTable#hash_key}.
	// Experimental.
	HashKey *string `field:"optional" json:"hashKey" yaml:"hashKey"`
	// key_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#key_schema AwsTable#key_schema}
	// Experimental.
	KeySchema interface{} `field:"optional" json:"keySchema" yaml:"keySchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#non_key_attributes AwsTable#non_key_attributes}.
	// Experimental.
	NonKeyAttributes *[]*string `field:"optional" json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
	// on_demand_throughput block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#on_demand_throughput AwsTable#on_demand_throughput}
	// Experimental.
	OnDemandThroughput *AwsTable_GlobalSecondaryIndexOnDemandThroughputProperty `field:"optional" json:"onDemandThroughput" yaml:"onDemandThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#range_key AwsTable#range_key}.
	// Experimental.
	RangeKey *string `field:"optional" json:"rangeKey" yaml:"rangeKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#read_capacity AwsTable#read_capacity}.
	// Experimental.
	ReadCapacity *float64 `field:"optional" json:"readCapacity" yaml:"readCapacity"`
	// warm_throughput block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#warm_throughput AwsTable#warm_throughput}
	// Experimental.
	WarmThroughput *AwsTable_GlobalSecondaryIndexWarmThroughputProperty `field:"optional" json:"warmThroughput" yaml:"warmThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#write_capacity AwsTable#write_capacity}.
	// Experimental.
	WriteCapacity *float64 `field:"optional" json:"writeCapacity" yaml:"writeCapacity"`
}

