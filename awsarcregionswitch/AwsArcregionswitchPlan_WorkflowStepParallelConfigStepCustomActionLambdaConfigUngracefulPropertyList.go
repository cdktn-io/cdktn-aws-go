package awsarcregionswitch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsarcregionswitch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsarcregionswitch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Get(index *float64) AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList
type jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList {
	_init_.Initialize()

	if err := validateNewAwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.AwsArcregionswitchPlan.WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList_Override(a AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.AwsArcregionswitchPlan.WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList) Get(index *float64) AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

