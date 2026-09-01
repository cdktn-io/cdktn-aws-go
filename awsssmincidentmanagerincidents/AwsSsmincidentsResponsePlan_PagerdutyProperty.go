package awsssmincidentmanagerincidents


// Experimental.
type AwsSsmincidentsResponsePlan_PagerdutyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#name AwsSsmincidentsResponsePlan#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#secret_id AwsSsmincidentsResponsePlan#secret_id}.
	// Experimental.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#service_id AwsSsmincidentsResponsePlan#service_id}.
	// Experimental.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
}

