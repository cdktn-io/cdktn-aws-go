package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference interface {
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
	IdleTimeoutInMinutes() *float64
	// Experimental.
	SetIdleTimeoutInMinutes(val *float64)
	// Experimental.
	IdleTimeoutInMinutesInput() *float64
	// Experimental.
	InternalValue() *AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty)
	// Experimental.
	LifecycleManagement() *string
	// Experimental.
	SetLifecycleManagement(val *string)
	// Experimental.
	LifecycleManagementInput() *string
	// Experimental.
	MaxIdleTimeoutInMinutes() *float64
	// Experimental.
	SetMaxIdleTimeoutInMinutes(val *float64)
	// Experimental.
	MaxIdleTimeoutInMinutesInput() *float64
	// Experimental.
	MinIdleTimeoutInMinutes() *float64
	// Experimental.
	SetMinIdleTimeoutInMinutes(val *float64)
	// Experimental.
	MinIdleTimeoutInMinutesInput() *float64
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
	ResetIdleTimeoutInMinutes()
	// Experimental.
	ResetLifecycleManagement()
	// Experimental.
	ResetMaxIdleTimeoutInMinutes()
	// Experimental.
	ResetMinIdleTimeoutInMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference
type jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) IdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) IdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) InternalValue() *AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty {
	var returns *AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) LifecycleManagement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) LifecycleManagementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) MaxIdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) MaxIdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) MinIdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) MinIdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference_Override(a AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetIdleTimeoutInMinutes(val *float64) {
	if err := j.validateSetIdleTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetInternalValue(val *AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetLifecycleManagement(val *string) {
	if err := j.validateSetLifecycleManagementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleManagement",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetMaxIdleTimeoutInMinutes(val *float64) {
	if err := j.validateSetMaxIdleTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetMinIdleTimeoutInMinutes(val *float64) {
	if err := j.validateSetMinIdleTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIdleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ResetIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ResetLifecycleManagement() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycleManagement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ResetMaxIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ResetMinIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetMinIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

