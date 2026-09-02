package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssecurityhub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssecurityhub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference interface {
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
	EnabledStandardArns() *[]*string
	// Experimental.
	SetEnabledStandardArns(val *[]*string)
	// Experimental.
	EnabledStandardArnsInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfConfigurationPolicy_ConfigurationPolicyProperty
	// Experimental.
	SetInternalValue(val *TfConfigurationPolicy_ConfigurationPolicyProperty)
	// Experimental.
	SecurityControlsConfiguration() TfConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference
	// Experimental.
	SecurityControlsConfigurationInput() *TfConfigurationPolicy_SecurityControlsConfigurationProperty
	// Experimental.
	ServiceEnabled() interface{}
	// Experimental.
	SetServiceEnabled(val interface{})
	// Experimental.
	ServiceEnabledInput() interface{}
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
	PutSecurityControlsConfiguration(value *TfConfigurationPolicy_SecurityControlsConfigurationProperty)
	// Experimental.
	ResetEnabledStandardArns()
	// Experimental.
	ResetSecurityControlsConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference
type jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) EnabledStandardArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledStandardArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) EnabledStandardArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledStandardArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) InternalValue() *TfConfigurationPolicy_ConfigurationPolicyProperty {
	var returns *TfConfigurationPolicy_ConfigurationPolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) SecurityControlsConfiguration() TfConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference {
	var returns TfConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"securityControlsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) SecurityControlsConfigurationInput() *TfConfigurationPolicy_SecurityControlsConfigurationProperty {
	var returns *TfConfigurationPolicy_SecurityControlsConfigurationProperty
	_jsii_.Get(
		j,
		"securityControlsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ServiceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ServiceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConfigurationPolicy_ConfigurationPolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-hub.TfConfigurationPolicy.ConfigurationPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference_Override(t TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-hub.TfConfigurationPolicy.ConfigurationPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference)SetEnabledStandardArns(val *[]*string) {
	if err := j.validateSetEnabledStandardArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledStandardArns",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference)SetInternalValue(val *TfConfigurationPolicy_ConfigurationPolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference)SetServiceEnabled(val interface{}) {
	if err := j.validateSetServiceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceEnabled",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) PutSecurityControlsConfiguration(value *TfConfigurationPolicy_SecurityControlsConfigurationProperty) {
	if err := t.validatePutSecurityControlsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecurityControlsConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ResetEnabledStandardArns() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabledStandardArns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ResetSecurityControlsConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityControlsConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

