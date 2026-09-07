package bedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	HistoricalContextWindowSize() *float64
	// Experimental.
	SetHistoricalContextWindowSize(val *float64)
	// Experimental.
	HistoricalContextWindowSizeInput() *float64
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	InvocationConfiguration() AwsMemoryStrategy_InvocationConfigurationPropertyList
	// Experimental.
	InvocationConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TriggerConditions() AwsMemoryStrategy_TriggerConditionsPropertyList
	// Experimental.
	TriggerConditionsActual() AwsMemoryStrategy_TriggerConditionsActualPropertyList
	// Experimental.
	TriggerConditionsInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	PutInvocationConfiguration(value interface{})
	// Experimental.
	PutTriggerConditions(value interface{})
	// Experimental.
	ResetHistoricalContextWindowSize()
	// Experimental.
	ResetInvocationConfiguration()
	// Experimental.
	ResetTriggerConditions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference
type jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) HistoricalContextWindowSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"historicalContextWindowSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) HistoricalContextWindowSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"historicalContextWindowSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) InvocationConfiguration() AwsMemoryStrategy_InvocationConfigurationPropertyList {
	var returns AwsMemoryStrategy_InvocationConfigurationPropertyList
	_jsii_.Get(
		j,
		"invocationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) InvocationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invocationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) TriggerConditions() AwsMemoryStrategy_TriggerConditionsPropertyList {
	var returns AwsMemoryStrategy_TriggerConditionsPropertyList
	_jsii_.Get(
		j,
		"triggerConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) TriggerConditionsActual() AwsMemoryStrategy_TriggerConditionsActualPropertyList {
	var returns AwsMemoryStrategy_TriggerConditionsActualPropertyList
	_jsii_.Get(
		j,
		"triggerConditionsActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) TriggerConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerConditionsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsMemoryStrategy.SelfManagedConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference_Override(a AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsMemoryStrategy.SelfManagedConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference)SetHistoricalContextWindowSize(val *float64) {
	if err := j.validateSetHistoricalContextWindowSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"historicalContextWindowSize",
		val,
	)
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) PutInvocationConfiguration(value interface{}) {
	if err := a.validatePutInvocationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInvocationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) PutTriggerConditions(value interface{}) {
	if err := a.validatePutTriggerConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTriggerConditions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) ResetHistoricalContextWindowSize() {
	_jsii_.InvokeVoid(
		a,
		"resetHistoricalContextWindowSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) ResetInvocationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetInvocationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) ResetTriggerConditions() {
	_jsii_.InvokeVoid(
		a,
		"resetTriggerConditions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMemoryStrategy_SelfManagedConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

