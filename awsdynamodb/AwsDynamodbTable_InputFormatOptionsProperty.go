package awsdynamodb


// Experimental.
type AwsDynamodbTable_InputFormatOptionsProperty struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#csv AwsDynamodbTable#csv}
	// Experimental.
	Csv *AwsDynamodbTable_CsvProperty `field:"optional" json:"csv" yaml:"csv"`
}

