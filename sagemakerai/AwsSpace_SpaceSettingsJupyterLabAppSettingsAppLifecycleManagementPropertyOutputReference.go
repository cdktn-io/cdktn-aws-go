package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference interface {
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
	IdleSettings() AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference
	// Experimental.
	IdleSettingsInput() *AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsProperty
	// Experimental.
	InternalValue() *AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementProperty
	// Experimental.
	SetInternalValue(val *AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementProperty)
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
	PutIdleSettings(value *AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsProperty)
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

// The jsii proxy struct for AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference
type jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) IdleSettings() AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference {
	var returns AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"idleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) IdleSettingsInput() *AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsProperty {
	var returns *AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsProperty
	_jsii_.Get(
		j,
		"idleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) InternalValue() *AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementProperty {
	var returns *AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSpace.SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference_Override(a AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSpace.SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference)SetInternalValue(val *AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) PutIdleSettings(value *AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsProperty) {
	if err := a.validatePutIdleSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIdleSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) ResetIdleSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetIdleSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

