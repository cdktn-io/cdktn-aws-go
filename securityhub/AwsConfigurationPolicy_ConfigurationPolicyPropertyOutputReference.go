package securityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/securityhub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/securityhub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference interface {
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
	InternalValue() *AwsConfigurationPolicy_ConfigurationPolicyProperty
	// Experimental.
	SetInternalValue(val *AwsConfigurationPolicy_ConfigurationPolicyProperty)
	// Experimental.
	SecurityControlsConfiguration() AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference
	// Experimental.
	SecurityControlsConfigurationInput() *AwsConfigurationPolicy_SecurityControlsConfigurationProperty
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
	PutSecurityControlsConfiguration(value *AwsConfigurationPolicy_SecurityControlsConfigurationProperty)
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

// The jsii proxy struct for AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference
type jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) EnabledStandardArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledStandardArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) EnabledStandardArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledStandardArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) InternalValue() *AwsConfigurationPolicy_ConfigurationPolicyProperty {
	var returns *AwsConfigurationPolicy_ConfigurationPolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) SecurityControlsConfiguration() AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference {
	var returns AwsConfigurationPolicy_SecurityControlsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"securityControlsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) SecurityControlsConfigurationInput() *AwsConfigurationPolicy_SecurityControlsConfigurationProperty {
	var returns *AwsConfigurationPolicy_SecurityControlsConfigurationProperty
	_jsii_.Get(
		j,
		"securityControlsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ServiceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ServiceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsConfigurationPolicy.ConfigurationPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference_Override(a AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsConfigurationPolicy.ConfigurationPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference)SetEnabledStandardArns(val *[]*string) {
	if err := j.validateSetEnabledStandardArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledStandardArns",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference)SetInternalValue(val *AwsConfigurationPolicy_ConfigurationPolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference)SetServiceEnabled(val interface{}) {
	if err := j.validateSetServiceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) PutSecurityControlsConfiguration(value *AwsConfigurationPolicy_SecurityControlsConfigurationProperty) {
	if err := a.validatePutSecurityControlsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecurityControlsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ResetEnabledStandardArns() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabledStandardArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ResetSecurityControlsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityControlsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ConfigurationPolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

