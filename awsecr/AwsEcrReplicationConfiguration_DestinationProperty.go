package awsecr


// Experimental.
type AwsEcrReplicationConfiguration_DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_replication_configuration#region AwsEcrReplicationConfiguration#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_replication_configuration#registry_id AwsEcrReplicationConfiguration#registry_id}.
	// Experimental.
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
}

