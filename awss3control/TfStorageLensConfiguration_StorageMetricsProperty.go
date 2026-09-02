package awss3control


// Experimental.
type TfStorageLensConfiguration_StorageMetricsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#enabled TfStorageLensConfiguration#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// selection_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#selection_criteria TfStorageLensConfiguration#selection_criteria}
	// Experimental.
	SelectionCriteria *TfStorageLensConfiguration_SelectionCriteriaProperty `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
}

