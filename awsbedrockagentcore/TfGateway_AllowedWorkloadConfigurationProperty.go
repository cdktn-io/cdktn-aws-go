package awsbedrockagentcore


// Experimental.
type TfGateway_AllowedWorkloadConfigurationProperty struct {
	// hosting_environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#hosting_environment TfGateway#hosting_environment}
	// Experimental.
	HostingEnvironment interface{} `field:"optional" json:"hostingEnvironment" yaml:"hostingEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#workload_identities TfGateway#workload_identities}.
	// Experimental.
	WorkloadIdentities *[]*string `field:"optional" json:"workloadIdentities" yaml:"workloadIdentities"`
}

