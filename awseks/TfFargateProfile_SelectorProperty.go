package awseks


// Experimental.
type TfFargateProfile_SelectorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_fargate_profile#namespace TfFargateProfile#namespace}.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_fargate_profile#labels TfFargateProfile#labels}.
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

