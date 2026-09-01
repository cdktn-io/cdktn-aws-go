package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference interface {
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
	InternalValue() *AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecProperty)
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

// The jsii proxy struct for AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference
type jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) InternalValue() *AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecProperty {
	var returns *AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) LifecycleConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) LifecycleConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageVersionAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageVersionAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageVersionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerSpace.SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference_Override(a AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerSpace.SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference)SetInternalValue(val *AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference)SetLifecycleConfigArn(val *string) {
	if err := j.validateSetLifecycleConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArn",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference)SetSagemakerImageArn(val *string) {
	if err := j.validateSetSagemakerImageArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sagemakerImageArn",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference)SetSagemakerImageVersionAlias(val *string) {
	if err := j.validateSetSagemakerImageVersionAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sagemakerImageVersionAlias",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference)SetSagemakerImageVersionArn(val *string) {
	if err := j.validateSetSagemakerImageVersionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sagemakerImageVersionArn",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) ResetLifecycleConfigArn() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycleConfigArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) ResetSagemakerImageArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSagemakerImageArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) ResetSagemakerImageVersionAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetSagemakerImageVersionAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) ResetSagemakerImageVersionArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSagemakerImageVersionArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

