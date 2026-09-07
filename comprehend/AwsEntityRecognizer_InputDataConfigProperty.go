package comprehend


// Experimental.
type AwsEntityRecognizer_InputDataConfigProperty struct {
	// entity_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#entity_types AwsEntityRecognizer#entity_types}
	// Experimental.
	EntityTypes interface{} `field:"required" json:"entityTypes" yaml:"entityTypes"`
	// annotations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#annotations AwsEntityRecognizer#annotations}
	// Experimental.
	Annotations *AwsEntityRecognizer_AnnotationsProperty `field:"optional" json:"annotations" yaml:"annotations"`
	// augmented_manifests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#augmented_manifests AwsEntityRecognizer#augmented_manifests}
	// Experimental.
	AugmentedManifests interface{} `field:"optional" json:"augmentedManifests" yaml:"augmentedManifests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#data_format AwsEntityRecognizer#data_format}.
	// Experimental.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// documents block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#documents AwsEntityRecognizer#documents}
	// Experimental.
	Documents *AwsEntityRecognizer_DocumentsProperty `field:"optional" json:"documents" yaml:"documents"`
	// entity_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#entity_list AwsEntityRecognizer#entity_list}
	// Experimental.
	EntityList *AwsEntityRecognizer_EntityListProperty `field:"optional" json:"entityList" yaml:"entityList"`
}

