package ssmincidentmanagerincidents


// Experimental.
type AwsResponsePlan_PagerdutyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#name AwsResponsePlan#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#secret_id AwsResponsePlan#secret_id}.
	// Experimental.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#service_id AwsResponsePlan#service_id}.
	// Experimental.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
}

