package dynamodb


// Experimental.
type AwsTable_InputFormatOptionsProperty struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#csv AwsTable#csv}
	// Experimental.
	Csv *AwsTable_CsvProperty `field:"optional" json:"csv" yaml:"csv"`
}

