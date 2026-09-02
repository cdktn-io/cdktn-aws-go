package awssagemakerai


// Experimental.
type TfProject_ServiceCatalogProvisioningDetailsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_project#product_id TfProject#product_id}.
	// Experimental.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_project#path_id TfProject#path_id}.
	// Experimental.
	PathId *string `field:"optional" json:"pathId" yaml:"pathId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_project#provisioning_artifact_id TfProject#provisioning_artifact_id}.
	// Experimental.
	ProvisioningArtifactId *string `field:"optional" json:"provisioningArtifactId" yaml:"provisioningArtifactId"`
	// provisioning_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_project#provisioning_parameter TfProject#provisioning_parameter}
	// Experimental.
	ProvisioningParameter interface{} `field:"optional" json:"provisioningParameter" yaml:"provisioningParameter"`
}

