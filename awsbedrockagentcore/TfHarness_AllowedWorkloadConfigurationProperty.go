package awsbedrockagentcore


// Experimental.
type TfHarness_AllowedWorkloadConfigurationProperty struct {
	// hosting_environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#hosting_environment TfHarness#hosting_environment}
	// Experimental.
	HostingEnvironment interface{} `field:"optional" json:"hostingEnvironment" yaml:"hostingEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#workload_identities TfHarness#workload_identities}.
	// Experimental.
	WorkloadIdentities *[]*string `field:"optional" json:"workloadIdentities" yaml:"workloadIdentities"`
}

