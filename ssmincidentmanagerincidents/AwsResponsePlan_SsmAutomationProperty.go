package ssmincidentmanagerincidents


// Experimental.
type AwsResponsePlan_SsmAutomationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#document_name AwsResponsePlan#document_name}.
	// Experimental.
	DocumentName *string `field:"required" json:"documentName" yaml:"documentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#role_arn AwsResponsePlan#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#document_version AwsResponsePlan#document_version}.
	// Experimental.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#dynamic_parameters AwsResponsePlan#dynamic_parameters}.
	// Experimental.
	DynamicParameters *map[string]*string `field:"optional" json:"dynamicParameters" yaml:"dynamicParameters"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#parameter AwsResponsePlan#parameter}
	// Experimental.
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#target_account AwsResponsePlan#target_account}.
	// Experimental.
	TargetAccount *string `field:"optional" json:"targetAccount" yaml:"targetAccount"`
}

