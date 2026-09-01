package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference interface {
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
	IdleSettings() AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference
	// Experimental.
	IdleSettingsInput() *AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsProperty
	// Experimental.
	InternalValue() *AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty)
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
	PutIdleSettings(value *AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsProperty)
	// Experimental.
	ResetIdleSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference
type jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) IdleSettings() AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference {
	var returns AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"idleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) IdleSettingsInput() *AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsProperty {
	var returns *AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsProperty
	_jsii_.Get(
		j,
		"idleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) InternalValue() *AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty {
	var returns *AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerUserProfile.UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference_Override(a AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerUserProfile.UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference)SetInternalValue(val *AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) PutIdleSettings(value *AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsProperty) {
	if err := a.validatePutIdleSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIdleSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) ResetIdleSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetIdleSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

