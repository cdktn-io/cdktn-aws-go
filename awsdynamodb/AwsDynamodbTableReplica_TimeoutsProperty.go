package awsdynamodb


// Experimental.
type AwsDynamodbTableReplica_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_replica#create AwsDynamodbTableReplica#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_replica#delete AwsDynamodbTableReplica#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_replica#update AwsDynamodbTableReplica#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

