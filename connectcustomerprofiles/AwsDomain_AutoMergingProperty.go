package connectcustomerprofiles


// Experimental.
type AwsDomain_AutoMergingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#enabled AwsDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// conflict_resolution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#conflict_resolution AwsDomain#conflict_resolution}
	// Experimental.
	ConflictResolution *AwsDomain_MatchingAutoMergingConflictResolutionProperty `field:"optional" json:"conflictResolution" yaml:"conflictResolution"`
	// consolidation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#consolidation AwsDomain#consolidation}
	// Experimental.
	Consolidation *AwsDomain_ConsolidationProperty `field:"optional" json:"consolidation" yaml:"consolidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#min_allowed_confidence_score_for_merging AwsDomain#min_allowed_confidence_score_for_merging}.
	// Experimental.
	MinAllowedConfidenceScoreForMerging *float64 `field:"optional" json:"minAllowedConfidenceScoreForMerging" yaml:"minAllowedConfidenceScoreForMerging"`
}

