package awsssmincidentmanagerincidents


// Experimental.
type TfResponsePlan_SsmAutomationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#document_name TfResponsePlan#document_name}.
	// Experimental.
	DocumentName *string `field:"required" json:"documentName" yaml:"documentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#role_arn TfResponsePlan#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#document_version TfResponsePlan#document_version}.
	// Experimental.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#dynamic_parameters TfResponsePlan#dynamic_parameters}.
	// Experimental.
	DynamicParameters *map[string]*string `field:"optional" json:"dynamicParameters" yaml:"dynamicParameters"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#parameter TfResponsePlan#parameter}
	// Experimental.
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#target_account TfResponsePlan#target_account}.
	// Experimental.
	TargetAccount *string `field:"optional" json:"targetAccount" yaml:"targetAccount"`
}

