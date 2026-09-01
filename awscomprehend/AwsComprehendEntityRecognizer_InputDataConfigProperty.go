package awscomprehend


// Experimental.
type AwsComprehendEntityRecognizer_InputDataConfigProperty struct {
	// entity_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#entity_types AwsComprehendEntityRecognizer#entity_types}
	// Experimental.
	EntityTypes interface{} `field:"required" json:"entityTypes" yaml:"entityTypes"`
	// annotations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#annotations AwsComprehendEntityRecognizer#annotations}
	// Experimental.
	Annotations *AwsComprehendEntityRecognizer_AnnotationsProperty `field:"optional" json:"annotations" yaml:"annotations"`
	// augmented_manifests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#augmented_manifests AwsComprehendEntityRecognizer#augmented_manifests}
	// Experimental.
	AugmentedManifests interface{} `field:"optional" json:"augmentedManifests" yaml:"augmentedManifests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#data_format AwsComprehendEntityRecognizer#data_format}.
	// Experimental.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// documents block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#documents AwsComprehendEntityRecognizer#documents}
	// Experimental.
	Documents *AwsComprehendEntityRecognizer_DocumentsProperty `field:"optional" json:"documents" yaml:"documents"`
	// entity_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#entity_list AwsComprehendEntityRecognizer#entity_list}
	// Experimental.
	EntityList *AwsComprehendEntityRecognizer_EntityListProperty `field:"optional" json:"entityList" yaml:"entityList"`
}

