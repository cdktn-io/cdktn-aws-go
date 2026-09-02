package awsappflow


// Experimental.
type TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#prefix_format TfFlow#prefix_format}.
	// Experimental.
	PrefixFormat *string `field:"optional" json:"prefixFormat" yaml:"prefixFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#prefix_hierarchy TfFlow#prefix_hierarchy}.
	// Experimental.
	PrefixHierarchy *[]*string `field:"optional" json:"prefixHierarchy" yaml:"prefixHierarchy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#prefix_type TfFlow#prefix_type}.
	// Experimental.
	PrefixType *string `field:"optional" json:"prefixType" yaml:"prefixType"`
}

