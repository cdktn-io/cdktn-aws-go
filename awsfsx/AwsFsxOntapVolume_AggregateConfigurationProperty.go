package awsfsx


// Experimental.
type AwsFsxOntapVolume_AggregateConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#aggregates AwsFsxOntapVolume#aggregates}.
	// Experimental.
	Aggregates *[]*string `field:"optional" json:"aggregates" yaml:"aggregates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#constituents_per_aggregate AwsFsxOntapVolume#constituents_per_aggregate}.
	// Experimental.
	ConstituentsPerAggregate *float64 `field:"optional" json:"constituentsPerAggregate" yaml:"constituentsPerAggregate"`
}

