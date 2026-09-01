package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference interface {
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
	HighAction() AwsCognitoRiskConfiguration_HighActionPropertyOutputReference
	// Experimental.
	HighActionInput() *AwsCognitoRiskConfiguration_HighActionProperty
	// Experimental.
	InternalValue() *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty
	// Experimental.
	SetInternalValue(val *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty)
	// Experimental.
	LowAction() AwsCognitoRiskConfiguration_LowActionPropertyOutputReference
	// Experimental.
	LowActionInput() *AwsCognitoRiskConfiguration_LowActionProperty
	// Experimental.
	MediumAction() AwsCognitoRiskConfiguration_MediumActionPropertyOutputReference
	// Experimental.
	MediumActionInput() *AwsCognitoRiskConfiguration_MediumActionProperty
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
	PutHighAction(value *AwsCognitoRiskConfiguration_HighActionProperty)
	// Experimental.
	PutLowAction(value *AwsCognitoRiskConfiguration_LowActionProperty)
	// Experimental.
	PutMediumAction(value *AwsCognitoRiskConfiguration_MediumActionProperty)
	// Experimental.
	ResetHighAction()
	// Experimental.
	ResetLowAction()
	// Experimental.
	ResetMediumAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference
type jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) HighAction() AwsCognitoRiskConfiguration_HighActionPropertyOutputReference {
	var returns AwsCognitoRiskConfiguration_HighActionPropertyOutputReference
	_jsii_.Get(
		j,
		"highAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) HighActionInput() *AwsCognitoRiskConfiguration_HighActionProperty {
	var returns *AwsCognitoRiskConfiguration_HighActionProperty
	_jsii_.Get(
		j,
		"highActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) InternalValue() *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty {
	var returns *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) LowAction() AwsCognitoRiskConfiguration_LowActionPropertyOutputReference {
	var returns AwsCognitoRiskConfiguration_LowActionPropertyOutputReference
	_jsii_.Get(
		j,
		"lowAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) LowActionInput() *AwsCognitoRiskConfiguration_LowActionProperty {
	var returns *AwsCognitoRiskConfiguration_LowActionProperty
	_jsii_.Get(
		j,
		"lowActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) MediumAction() AwsCognitoRiskConfiguration_MediumActionPropertyOutputReference {
	var returns AwsCognitoRiskConfiguration_MediumActionPropertyOutputReference
	_jsii_.Get(
		j,
		"mediumAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) MediumActionInput() *AwsCognitoRiskConfiguration_MediumActionProperty {
	var returns *AwsCognitoRiskConfiguration_MediumActionProperty
	_jsii_.Get(
		j,
		"mediumActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoRiskConfiguration.AccountTakeoverRiskConfigurationActionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference_Override(a AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoRiskConfiguration.AccountTakeoverRiskConfigurationActionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference)SetInternalValue(val *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) PutHighAction(value *AwsCognitoRiskConfiguration_HighActionProperty) {
	if err := a.validatePutHighActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHighAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) PutLowAction(value *AwsCognitoRiskConfiguration_LowActionProperty) {
	if err := a.validatePutLowActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLowAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) PutMediumAction(value *AwsCognitoRiskConfiguration_MediumActionProperty) {
	if err := a.validatePutMediumActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMediumAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) ResetHighAction() {
	_jsii_.InvokeVoid(
		a,
		"resetHighAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) ResetLowAction() {
	_jsii_.InvokeVoid(
		a,
		"resetLowAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) ResetMediumAction() {
	_jsii_.InvokeVoid(
		a,
		"resetMediumAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

