package awsdynamodb


// Experimental.
type AwsDynamodbGlobalSecondaryIndex_OnDemandThroughputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_global_secondary_index#max_read_request_units AwsDynamodbGlobalSecondaryIndex#max_read_request_units}.
	// Experimental.
	MaxReadRequestUnits *float64 `field:"optional" json:"maxReadRequestUnits" yaml:"maxReadRequestUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_global_secondary_index#max_write_request_units AwsDynamodbGlobalSecondaryIndex#max_write_request_units}.
	// Experimental.
	MaxWriteRequestUnits *float64 `field:"optional" json:"maxWriteRequestUnits" yaml:"maxWriteRequestUnits"`
}

