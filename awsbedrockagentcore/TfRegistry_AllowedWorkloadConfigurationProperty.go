package awsbedrockagentcore


// Experimental.
type TfRegistry_AllowedWorkloadConfigurationProperty struct {
	// hosting_environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_registry#hosting_environment TfRegistry#hosting_environment}
	// Experimental.
	HostingEnvironment interface{} `field:"optional" json:"hostingEnvironment" yaml:"hostingEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_registry#workload_identities TfRegistry#workload_identities}.
	// Experimental.
	WorkloadIdentities *[]*string `field:"optional" json:"workloadIdentities" yaml:"workloadIdentities"`
}

