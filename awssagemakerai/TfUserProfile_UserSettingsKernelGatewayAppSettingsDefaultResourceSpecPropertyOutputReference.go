package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference interface {
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
	InstanceType() *string
	// Experimental.
	SetInstanceType(val *string)
	// Experimental.
	InstanceTypeInput() *string
	// Experimental.
	InternalValue() *TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty
	// Experimental.
	SetInternalValue(val *TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty)
	// Experimental.
	LifecycleConfigArn() *string
	// Experimental.
	SetLifecycleConfigArn(val *string)
	// Experimental.
	LifecycleConfigArnInput() *string
	// Experimental.
	SagemakerImageArn() *string
	// Experimental.
	SetSagemakerImageArn(val *string)
	// Experimental.
	SagemakerImageArnInput() *string
	// Experimental.
	SagemakerImageVersionAlias() *string
	// Experimental.
	SetSagemakerImageVersionAlias(val *string)
	// Experimental.
	SagemakerImageVersionAliasInput() *string
	// Experimental.
	SagemakerImageVersionArn() *string
	// Experimental.
	SetSagemakerImageVersionArn(val *string)
	// Experimental.
	SagemakerImageVersionArnInput() *string
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
	ResetInstanceType()
	// Experimental.
	ResetLifecycleConfigArn()
	// Experimental.
	ResetSagemakerImageArn()
	// Experimental.
	ResetSagemakerImageVersionAlias()
	// Experimental.
	ResetSagemakerImageVersionArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference
type jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) InternalValue() *TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty {
	var returns *TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) LifecycleConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) LifecycleConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageVersionAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageVersionAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageVersionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference_Override(t TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetInternalValue(val *TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetLifecycleConfigArn(val *string) {
	if err := j.validateSetLifecycleConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArn",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetSagemakerImageArn(val *string) {
	if err := j.validateSetSagemakerImageArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sagemakerImageArn",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetSagemakerImageVersionAlias(val *string) {
	if err := j.validateSetSagemakerImageVersionAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sagemakerImageVersionAlias",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetSagemakerImageVersionArn(val *string) {
	if err := j.validateSetSagemakerImageVersionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sagemakerImageVersionArn",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ResetLifecycleConfigArn() {
	_jsii_.InvokeVoid(
		t,
		"resetLifecycleConfigArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ResetSagemakerImageArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSagemakerImageArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ResetSagemakerImageVersionAlias() {
	_jsii_.InvokeVoid(
		t,
		"resetSagemakerImageVersionAlias",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ResetSagemakerImageVersionArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSagemakerImageVersionArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

