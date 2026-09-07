package cognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Actions() AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationActionsPropertyOutputReference
	// Experimental.
	ActionsInput() *AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationActionsProperty
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
	EventFilter() *[]*string
	// Experimental.
	SetEventFilter(val *[]*string)
	// Experimental.
	EventFilterInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationProperty)
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
	PutActions(value *AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationActionsProperty)
	// Experimental.
	ResetEventFilter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference
type jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) Actions() AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationActionsPropertyOutputReference {
	var returns AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationActionsPropertyOutputReference
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) ActionsInput() *AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationActionsProperty {
	var returns *AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationActionsProperty
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) EventFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) EventFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) InternalValue() *AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationProperty {
	var returns *AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsRiskConfiguration.CompromisedCredentialsRiskConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference_Override(a AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsRiskConfiguration.CompromisedCredentialsRiskConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference)SetEventFilter(val *[]*string) {
	if err := j.validateSetEventFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventFilter",
		val,
	)
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference)SetInternalValue(val *AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) PutActions(value *AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationActionsProperty) {
	if err := a.validatePutActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) ResetEventFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetEventFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRiskConfiguration_CompromisedCredentialsRiskConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

