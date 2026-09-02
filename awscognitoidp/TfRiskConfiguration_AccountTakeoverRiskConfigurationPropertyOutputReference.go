package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Actions() TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference
	// Experimental.
	ActionsInput() *TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty
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
	InternalValue() *TfRiskConfiguration_AccountTakeoverRiskConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfRiskConfiguration_AccountTakeoverRiskConfigurationProperty)
	// Experimental.
	NotifyConfiguration() TfRiskConfiguration_NotifyConfigurationPropertyOutputReference
	// Experimental.
	NotifyConfigurationInput() *TfRiskConfiguration_NotifyConfigurationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutActions(value *TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty)
	// Experimental.
	PutNotifyConfiguration(value *TfRiskConfiguration_NotifyConfigurationProperty)
	// Experimental.
	ResetNotifyConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference
type jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) Actions() TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference {
	var returns TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) ActionsInput() *TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty {
	var returns *TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) InternalValue() *TfRiskConfiguration_AccountTakeoverRiskConfigurationProperty {
	var returns *TfRiskConfiguration_AccountTakeoverRiskConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) NotifyConfiguration() TfRiskConfiguration_NotifyConfigurationPropertyOutputReference {
	var returns TfRiskConfiguration_NotifyConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"notifyConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) NotifyConfigurationInput() *TfRiskConfiguration_NotifyConfigurationProperty {
	var returns *TfRiskConfiguration_NotifyConfigurationProperty
	_jsii_.Get(
		j,
		"notifyConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfRiskConfiguration.AccountTakeoverRiskConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference_Override(t TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfRiskConfiguration.AccountTakeoverRiskConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference)SetInternalValue(val *TfRiskConfiguration_AccountTakeoverRiskConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) PutActions(value *TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty) {
	if err := t.validatePutActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putActions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) PutNotifyConfiguration(value *TfRiskConfiguration_NotifyConfigurationProperty) {
	if err := t.validatePutNotifyConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNotifyConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) ResetNotifyConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetNotifyConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

