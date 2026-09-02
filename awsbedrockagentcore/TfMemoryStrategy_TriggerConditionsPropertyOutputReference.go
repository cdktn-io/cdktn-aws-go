package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMemoryStrategy_TriggerConditionsPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MessageBasedTrigger() TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerPropertyList
	// Experimental.
	MessageBasedTriggerInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeBasedTrigger() TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsTimeBasedTriggerPropertyList
	// Experimental.
	TimeBasedTriggerInput() interface{}
	// Experimental.
	TokenBasedTrigger() TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsTokenBasedTriggerPropertyList
	// Experimental.
	TokenBasedTriggerInput() interface{}
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
	PutMessageBasedTrigger(value interface{})
	// Experimental.
	PutTimeBasedTrigger(value interface{})
	// Experimental.
	PutTokenBasedTrigger(value interface{})
	// Experimental.
	ResetMessageBasedTrigger()
	// Experimental.
	ResetTimeBasedTrigger()
	// Experimental.
	ResetTokenBasedTrigger()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfMemoryStrategy_TriggerConditionsPropertyOutputReference
type jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) MessageBasedTrigger() TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerPropertyList {
	var returns TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerPropertyList
	_jsii_.Get(
		j,
		"messageBasedTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) MessageBasedTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageBasedTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) TimeBasedTrigger() TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsTimeBasedTriggerPropertyList {
	var returns TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsTimeBasedTriggerPropertyList
	_jsii_.Get(
		j,
		"timeBasedTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) TimeBasedTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeBasedTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) TokenBasedTrigger() TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsTokenBasedTriggerPropertyList {
	var returns TfMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsTokenBasedTriggerPropertyList
	_jsii_.Get(
		j,
		"tokenBasedTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) TokenBasedTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenBasedTriggerInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMemoryStrategy_TriggerConditionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfMemoryStrategy_TriggerConditionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMemoryStrategy_TriggerConditionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfMemoryStrategy.TriggerConditionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMemoryStrategy_TriggerConditionsPropertyOutputReference_Override(t TfMemoryStrategy_TriggerConditionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfMemoryStrategy.TriggerConditionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) PutMessageBasedTrigger(value interface{}) {
	if err := t.validatePutMessageBasedTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMessageBasedTrigger",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) PutTimeBasedTrigger(value interface{}) {
	if err := t.validatePutTimeBasedTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeBasedTrigger",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) PutTokenBasedTrigger(value interface{}) {
	if err := t.validatePutTokenBasedTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTokenBasedTrigger",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) ResetMessageBasedTrigger() {
	_jsii_.InvokeVoid(
		t,
		"resetMessageBasedTrigger",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) ResetTimeBasedTrigger() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeBasedTrigger",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) ResetTokenBasedTrigger() {
	_jsii_.InvokeVoid(
		t,
		"resetTokenBasedTrigger",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMemoryStrategy_TriggerConditionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

