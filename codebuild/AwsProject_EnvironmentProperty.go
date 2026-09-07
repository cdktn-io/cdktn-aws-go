package codebuild


// Experimental.
type AwsProject_EnvironmentProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#compute_type AwsProject#compute_type}.
	// Experimental.
	ComputeType *string `field:"required" json:"computeType" yaml:"computeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#image AwsProject#image}.
	// Experimental.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#type AwsProject#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#certificate AwsProject#certificate}.
	// Experimental.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// docker_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#docker_server AwsProject#docker_server}
	// Experimental.
	DockerServer *AwsProject_DockerServerProperty `field:"optional" json:"dockerServer" yaml:"dockerServer"`
	// environment_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#environment_variable AwsProject#environment_variable}
	// Experimental.
	EnvironmentVariable interface{} `field:"optional" json:"environmentVariable" yaml:"environmentVariable"`
	// fleet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#fleet AwsProject#fleet}
	// Experimental.
	Fleet *AwsProject_FleetProperty `field:"optional" json:"fleet" yaml:"fleet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#host_kernel AwsProject#host_kernel}.
	// Experimental.
	HostKernel *string `field:"optional" json:"hostKernel" yaml:"hostKernel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#image_pull_credentials_type AwsProject#image_pull_credentials_type}.
	// Experimental.
	ImagePullCredentialsType *string `field:"optional" json:"imagePullCredentialsType" yaml:"imagePullCredentialsType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#privileged_mode AwsProject#privileged_mode}.
	// Experimental.
	PrivilegedMode interface{} `field:"optional" json:"privilegedMode" yaml:"privilegedMode"`
	// registry_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#registry_credential AwsProject#registry_credential}
	// Experimental.
	RegistryCredential *AwsProject_RegistryCredentialProperty `field:"optional" json:"registryCredential" yaml:"registryCredential"`
}

