package awsconnect


// Experimental.
type TfUserHierarchyStructure_HierarchyStructureProperty struct {
	// level_five block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_five TfUserHierarchyStructure#level_five}
	// Experimental.
	LevelFive *TfUserHierarchyStructure_LevelFiveProperty `field:"optional" json:"levelFive" yaml:"levelFive"`
	// level_four block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_four TfUserHierarchyStructure#level_four}
	// Experimental.
	LevelFour *TfUserHierarchyStructure_LevelFourProperty `field:"optional" json:"levelFour" yaml:"levelFour"`
	// level_one block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_one TfUserHierarchyStructure#level_one}
	// Experimental.
	LevelOne *TfUserHierarchyStructure_LevelOneProperty `field:"optional" json:"levelOne" yaml:"levelOne"`
	// level_three block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_three TfUserHierarchyStructure#level_three}
	// Experimental.
	LevelThree *TfUserHierarchyStructure_LevelThreeProperty `field:"optional" json:"levelThree" yaml:"levelThree"`
	// level_two block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user_hierarchy_structure#level_two TfUserHierarchyStructure#level_two}
	// Experimental.
	LevelTwo *TfUserHierarchyStructure_LevelTwoProperty `field:"optional" json:"levelTwo" yaml:"levelTwo"`
}

