package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference interface {
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
	CustomImage() TfUserProfile_UserSettingsKernelGatewayAppSettingsCustomImagePropertyList
	// Experimental.
	CustomImageInput() interface{}
	// Experimental.
	DefaultResourceSpec() TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference
	// Experimental.
	DefaultResourceSpecInput() *TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfUserProfile_KernelGatewayAppSettingsProperty
	// Experimental.
	SetInternalValue(val *TfUserProfile_KernelGatewayAppSettingsProperty)
	// Experimental.
	LifecycleConfigArns() *[]*string
	// Experimental.
	SetLifecycleConfigArns(val *[]*string)
	// Experimental.
	LifecycleConfigArnsInput() *[]*string
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
	PutCustomImage(value interface{})
	// Experimental.
	PutDefaultResourceSpec(value *TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty)
	// Experimental.
	ResetCustomImage()
	// Experimental.
	ResetDefaultResourceSpec()
	// Experimental.
	ResetLifecycleConfigArns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference
type jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) CustomImage() TfUserProfile_UserSettingsKernelGatewayAppSettingsCustomImagePropertyList {
	var returns TfUserProfile_UserSettingsKernelGatewayAppSettingsCustomImagePropertyList
	_jsii_.Get(
		j,
		"customImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) CustomImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) DefaultResourceSpec() TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference {
	var returns TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) DefaultResourceSpecInput() *TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty {
	var returns *TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) InternalValue() *TfUserProfile_KernelGatewayAppSettingsProperty {
	var returns *TfUserProfile_KernelGatewayAppSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) LifecycleConfigArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) LifecycleConfigArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUserProfile_KernelGatewayAppSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUserProfile_KernelGatewayAppSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.KernelGatewayAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUserProfile_KernelGatewayAppSettingsPropertyOutputReference_Override(t TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.KernelGatewayAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference)SetInternalValue(val *TfUserProfile_KernelGatewayAppSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference)SetLifecycleConfigArns(val *[]*string) {
	if err := j.validateSetLifecycleConfigArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArns",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) PutCustomImage(value interface{}) {
	if err := t.validatePutCustomImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomImage",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) PutDefaultResourceSpec(value *TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty) {
	if err := t.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) ResetCustomImage() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomImage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) ResetLifecycleConfigArns() {
	_jsii_.InvokeVoid(
		t,
		"resetLifecycleConfigArns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

