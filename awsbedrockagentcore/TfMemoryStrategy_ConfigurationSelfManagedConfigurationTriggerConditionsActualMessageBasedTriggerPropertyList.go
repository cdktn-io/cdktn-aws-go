package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList
type jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList {
	_init_.Initialize()

	if err := validateNewTfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfMemoryStrategy.ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList_Override(t TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfMemoryStrategy.ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		t,
	)
}

func (j *jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (t *jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := t.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		t,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList) Get(index *float64) TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyOutputReference {
	if err := t.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyOutputReference

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsActualMessageBasedTriggerPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

