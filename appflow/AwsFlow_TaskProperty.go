package appflow


// Experimental.
type AwsFlow_TaskProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#task_type AwsFlow#task_type}.
	// Experimental.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// connector_operator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#connector_operator AwsFlow#connector_operator}
	// Experimental.
	ConnectorOperator interface{} `field:"optional" json:"connectorOperator" yaml:"connectorOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#destination_field AwsFlow#destination_field}.
	// Experimental.
	DestinationField *string `field:"optional" json:"destinationField" yaml:"destinationField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#source_fields AwsFlow#source_fields}.
	// Experimental.
	SourceFields *[]*string `field:"optional" json:"sourceFields" yaml:"sourceFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#task_properties AwsFlow#task_properties}.
	// Experimental.
	TaskProperties *map[string]*string `field:"optional" json:"taskProperties" yaml:"taskProperties"`
}

