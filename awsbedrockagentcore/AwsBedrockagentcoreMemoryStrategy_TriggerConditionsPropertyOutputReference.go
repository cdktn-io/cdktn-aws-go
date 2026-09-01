package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference interface {
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
	MessageBasedTrigger() AwsBedrockagentcoreMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerPropertyList
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
	TimeBasedTrigger() AwsBedrockagentcoreMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsTimeBasedTriggerPropertyList
	// Experimental.
	TimeBasedTriggerInput() interface{}
	// Experimental.
	TokenBasedTrigger() AwsBedrockagentcoreMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsTokenBasedTriggerPropertyList
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

// The jsii proxy struct for AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference
type jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) MessageBasedTrigger() AwsBedrockagentcoreMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerPropertyList {
	var returns AwsBedrockagentcoreMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerPropertyList
	_jsii_.Get(
		j,
		"messageBasedTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) MessageBasedTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageBasedTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) TimeBasedTrigger() AwsBedrockagentcoreMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsTimeBasedTriggerPropertyList {
	var returns AwsBedrockagentcoreMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsTimeBasedTriggerPropertyList
	_jsii_.Get(
		j,
		"timeBasedTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) TimeBasedTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeBasedTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) TokenBasedTrigger() AwsBedrockagentcoreMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsTokenBasedTriggerPropertyList {
	var returns AwsBedrockagentcoreMemoryStrategy_ConfigurationSelfManagedConfigurationTriggerConditionsTokenBasedTriggerPropertyList
	_jsii_.Get(
		j,
		"tokenBasedTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) TokenBasedTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenBasedTriggerInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreMemoryStrategy.TriggerConditionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference_Override(a AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreMemoryStrategy.TriggerConditionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) PutMessageBasedTrigger(value interface{}) {
	if err := a.validatePutMessageBasedTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMessageBasedTrigger",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) PutTimeBasedTrigger(value interface{}) {
	if err := a.validatePutTimeBasedTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeBasedTrigger",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) PutTokenBasedTrigger(value interface{}) {
	if err := a.validatePutTokenBasedTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTokenBasedTrigger",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) ResetMessageBasedTrigger() {
	_jsii_.InvokeVoid(
		a,
		"resetMessageBasedTrigger",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) ResetTimeBasedTrigger() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeBasedTrigger",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) ResetTokenBasedTrigger() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenBasedTrigger",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBedrockagentcoreMemoryStrategy_TriggerConditionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

