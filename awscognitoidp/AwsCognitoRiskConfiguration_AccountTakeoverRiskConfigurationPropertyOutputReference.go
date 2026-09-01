package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Actions() AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference
	// Experimental.
	ActionsInput() *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty
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
	InternalValue() *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationProperty)
	// Experimental.
	NotifyConfiguration() AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference
	// Experimental.
	NotifyConfigurationInput() *AwsCognitoRiskConfiguration_NotifyConfigurationProperty
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
	PutActions(value *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty)
	// Experimental.
	PutNotifyConfiguration(value *AwsCognitoRiskConfiguration_NotifyConfigurationProperty)
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

// The jsii proxy struct for AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference
type jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) Actions() AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference {
	var returns AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) ActionsInput() *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty {
	var returns *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) InternalValue() *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationProperty {
	var returns *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) NotifyConfiguration() AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference {
	var returns AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"notifyConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) NotifyConfigurationInput() *AwsCognitoRiskConfiguration_NotifyConfigurationProperty {
	var returns *AwsCognitoRiskConfiguration_NotifyConfigurationProperty
	_jsii_.Get(
		j,
		"notifyConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoRiskConfiguration.AccountTakeoverRiskConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference_Override(a AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoRiskConfiguration.AccountTakeoverRiskConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference)SetInternalValue(val *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) PutActions(value *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty) {
	if err := a.validatePutActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) PutNotifyConfiguration(value *AwsCognitoRiskConfiguration_NotifyConfigurationProperty) {
	if err := a.validatePutNotifyConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNotifyConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) ResetNotifyConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetNotifyConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

