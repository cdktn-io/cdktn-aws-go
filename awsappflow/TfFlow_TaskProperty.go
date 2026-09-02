package awsappflow


// Experimental.
type TfFlow_TaskProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#task_type TfFlow#task_type}.
	// Experimental.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// connector_operator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#connector_operator TfFlow#connector_operator}
	// Experimental.
	ConnectorOperator interface{} `field:"optional" json:"connectorOperator" yaml:"connectorOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#destination_field TfFlow#destination_field}.
	// Experimental.
	DestinationField *string `field:"optional" json:"destinationField" yaml:"destinationField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#source_fields TfFlow#source_fields}.
	// Experimental.
	SourceFields *[]*string `field:"optional" json:"sourceFields" yaml:"sourceFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#task_properties TfFlow#task_properties}.
	// Experimental.
	TaskProperties *map[string]*string `field:"optional" json:"taskProperties" yaml:"taskProperties"`
}

