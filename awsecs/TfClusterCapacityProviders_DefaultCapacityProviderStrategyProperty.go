package awsecs


// Experimental.
type TfClusterCapacityProviders_DefaultCapacityProviderStrategyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster_capacity_providers#capacity_provider TfClusterCapacityProviders#capacity_provider}.
	// Experimental.
	CapacityProvider *string `field:"required" json:"capacityProvider" yaml:"capacityProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster_capacity_providers#base TfClusterCapacityProviders#base}.
	// Experimental.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster_capacity_providers#weight TfClusterCapacityProviders#weight}.
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

