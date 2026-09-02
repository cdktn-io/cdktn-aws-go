package awscomprehend


// Experimental.
type TfEntityRecognizer_InputDataConfigProperty struct {
	// entity_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#entity_types TfEntityRecognizer#entity_types}
	// Experimental.
	EntityTypes interface{} `field:"required" json:"entityTypes" yaml:"entityTypes"`
	// annotations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#annotations TfEntityRecognizer#annotations}
	// Experimental.
	Annotations *TfEntityRecognizer_AnnotationsProperty `field:"optional" json:"annotations" yaml:"annotations"`
	// augmented_manifests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#augmented_manifests TfEntityRecognizer#augmented_manifests}
	// Experimental.
	AugmentedManifests interface{} `field:"optional" json:"augmentedManifests" yaml:"augmentedManifests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#data_format TfEntityRecognizer#data_format}.
	// Experimental.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// documents block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#documents TfEntityRecognizer#documents}
	// Experimental.
	Documents *TfEntityRecognizer_DocumentsProperty `field:"optional" json:"documents" yaml:"documents"`
	// entity_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#entity_list TfEntityRecognizer#entity_list}
	// Experimental.
	EntityList *TfEntityRecognizer_EntityListProperty `field:"optional" json:"entityList" yaml:"entityList"`
}

