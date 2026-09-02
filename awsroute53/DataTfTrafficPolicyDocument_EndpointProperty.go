package awsroute53


// Experimental.
type DataTfTrafficPolicyDocument_EndpointProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#id DataTfTrafficPolicyDocument#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#region DataTfTrafficPolicyDocument#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#type DataTfTrafficPolicyDocument#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#value DataTfTrafficPolicyDocument#value}.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

