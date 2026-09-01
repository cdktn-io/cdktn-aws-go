package awsdynamodb


// Experimental.
type AwsDynamodbTable_CsvProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#delimiter AwsDynamodbTable#delimiter}.
	// Experimental.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#header_list AwsDynamodbTable#header_list}.
	// Experimental.
	HeaderList *[]*string `field:"optional" json:"headerList" yaml:"headerList"`
}

