package awsconnectcustomerprofiles


// Experimental.
type TfDomain_AutoMergingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#enabled TfDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// conflict_resolution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#conflict_resolution TfDomain#conflict_resolution}
	// Experimental.
	ConflictResolution *TfDomain_MatchingAutoMergingConflictResolutionProperty `field:"optional" json:"conflictResolution" yaml:"conflictResolution"`
	// consolidation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#consolidation TfDomain#consolidation}
	// Experimental.
	Consolidation *TfDomain_ConsolidationProperty `field:"optional" json:"consolidation" yaml:"consolidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#min_allowed_confidence_score_for_merging TfDomain#min_allowed_confidence_score_for_merging}.
	// Experimental.
	MinAllowedConfidenceScoreForMerging *float64 `field:"optional" json:"minAllowedConfidenceScoreForMerging" yaml:"minAllowedConfidenceScoreForMerging"`
}

