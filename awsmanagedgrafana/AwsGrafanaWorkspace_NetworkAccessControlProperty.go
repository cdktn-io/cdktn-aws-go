package awsmanagedgrafana


// Experimental.
type AwsGrafanaWorkspace_NetworkAccessControlProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#prefix_list_ids AwsGrafanaWorkspace#prefix_list_ids}.
	// Experimental.
	PrefixListIds *[]*string `field:"required" json:"prefixListIds" yaml:"prefixListIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#vpce_ids AwsGrafanaWorkspace#vpce_ids}.
	// Experimental.
	VpceIds *[]*string `field:"required" json:"vpceIds" yaml:"vpceIds"`
}

