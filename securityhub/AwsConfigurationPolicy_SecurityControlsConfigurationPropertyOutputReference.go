package securityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/securityhub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/securityhub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference interface {
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
	DisabledControlIdentifiers() *[]*string
	// Experimental.
	SetDisabledControlIdentifiers(val *[]*string)
	// Experimental.
	DisabledControlIdentifiersInput() *[]*string
	// Experimental.
	EnabledControlIdentifiers() *[]*string
	// Experimental.
	SetEnabledControlIdentifiers(val *[]*string)
	// Experimental.
	EnabledControlIdentifiersInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsConfigurationPolicy_SecurityControlsConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsConfigurationPolicy_SecurityControlsConfigurationProperty)
	// Experimental.
	SecurityControlCustomParameter() AwsConfigurationPolicy_SecurityControlCustomParameterPropertyList
	// Experimental.
	SecurityControlCustomParameterInput() interface{}
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
	PutSecurityControlCustomParameter(value interface{})
	// Experimental.
	ResetDisabledControlIdentifiers()
	// Experimental.
	ResetEnabledControlIdentifiers()
	// Experimental.
	ResetSecurityControlCustomParameter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference
type jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) DisabledControlIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledControlIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) DisabledControlIdentifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledControlIdentifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) EnabledControlIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledControlIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) EnabledControlIdentifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledControlIdentifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) InternalValue() *AwsConfigurationPolicy_SecurityControlsConfigurationProperty {
	var returns *AwsConfigurationPolicy_SecurityControlsConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) SecurityControlCustomParameter() AwsConfigurationPolicy_SecurityControlCustomParameterPropertyList {
	var returns AwsConfigurationPolicy_SecurityControlCustomParameterPropertyList
	_jsii_.Get(
		j,
		"securityControlCustomParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) SecurityControlCustomParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityControlCustomParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsConfigurationPolicy.SecurityControlsConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference_Override(a AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsConfigurationPolicy.SecurityControlsConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference)SetDisabledControlIdentifiers(val *[]*string) {
	if err := j.validateSetDisabledControlIdentifiersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabledControlIdentifiers",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference)SetEnabledControlIdentifiers(val *[]*string) {
	if err := j.validateSetEnabledControlIdentifiersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledControlIdentifiers",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference)SetInternalValue(val *AwsConfigurationPolicy_SecurityControlsConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) PutSecurityControlCustomParameter(value interface{}) {
	if err := a.validatePutSecurityControlCustomParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecurityControlCustomParameter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) ResetDisabledControlIdentifiers() {
	_jsii_.InvokeVoid(
		a,
		"resetDisabledControlIdentifiers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) ResetEnabledControlIdentifiers() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabledControlIdentifiers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) ResetSecurityControlCustomParameter() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityControlCustomParameter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

