package connect


// Experimental.
type AwsUserHierarchyStructure_HierarchyStructureProperty struct {
	// level_five block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_five AwsUserHierarchyStructure#level_five}
	// Experimental.
	LevelFive *AwsUserHierarchyStructure_LevelFiveProperty `field:"optional" json:"levelFive" yaml:"levelFive"`
	// level_four block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_four AwsUserHierarchyStructure#level_four}
	// Experimental.
	LevelFour *AwsUserHierarchyStructure_LevelFourProperty `field:"optional" json:"levelFour" yaml:"levelFour"`
	// level_one block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_one AwsUserHierarchyStructure#level_one}
	// Experimental.
	LevelOne *AwsUserHierarchyStructure_LevelOneProperty `field:"optional" json:"levelOne" yaml:"levelOne"`
	// level_three block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_three AwsUserHierarchyStructure#level_three}
	// Experimental.
	LevelThree *AwsUserHierarchyStructure_LevelThreeProperty `field:"optional" json:"levelThree" yaml:"levelThree"`
	// level_two block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_two AwsUserHierarchyStructure#level_two}
	// Experimental.
	LevelTwo *AwsUserHierarchyStructure_LevelTwoProperty `field:"optional" json:"levelTwo" yaml:"levelTwo"`
}

