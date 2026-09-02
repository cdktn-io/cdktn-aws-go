package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference interface {
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
	CustomImage() TfDomain_DefaultUserSettingsKernelGatewayAppSettingsCustomImagePropertyList
	// Experimental.
	CustomImageInput() interface{}
	// Experimental.
	DefaultResourceSpec() TfDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference
	// Experimental.
	DefaultResourceSpecInput() *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty
	// Experimental.
	SetInternalValue(val *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty)
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
	PutDefaultResourceSpec(value *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty)
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

// The jsii proxy struct for TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference
type jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) CustomImage() TfDomain_DefaultUserSettingsKernelGatewayAppSettingsCustomImagePropertyList {
	var returns TfDomain_DefaultUserSettingsKernelGatewayAppSettingsCustomImagePropertyList
	_jsii_.Get(
		j,
		"customImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) CustomImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) DefaultResourceSpec() TfDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference {
	var returns TfDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) DefaultResourceSpecInput() *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty {
	var returns *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) InternalValue() *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty {
	var returns *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) LifecycleConfigArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) LifecycleConfigArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference_Override(t TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference)SetInternalValue(val *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference)SetLifecycleConfigArns(val *[]*string) {
	if err := j.validateSetLifecycleConfigArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArns",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) PutCustomImage(value interface{}) {
	if err := t.validatePutCustomImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomImage",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) PutDefaultResourceSpec(value *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty) {
	if err := t.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) ResetCustomImage() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomImage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) ResetLifecycleConfigArns() {
	_jsii_.InvokeVoid(
		t,
		"resetLifecycleConfigArns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

