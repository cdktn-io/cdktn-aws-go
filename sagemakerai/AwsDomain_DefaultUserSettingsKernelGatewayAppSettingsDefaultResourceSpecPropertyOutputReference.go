package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference interface {
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
	InternalValue() *AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty
	// Experimental.
	SetInternalValue(val *AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty)
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

// The jsii proxy struct for AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference
type jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) InternalValue() *AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty {
	var returns *AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) LifecycleConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) LifecycleConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageVersionAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageVersionAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) SagemakerImageVersionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerImageVersionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference_Override(a AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetInternalValue(val *AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetLifecycleConfigArn(val *string) {
	if err := j.validateSetLifecycleConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArn",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetSagemakerImageArn(val *string) {
	if err := j.validateSetSagemakerImageArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sagemakerImageArn",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetSagemakerImageVersionAlias(val *string) {
	if err := j.validateSetSagemakerImageVersionAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sagemakerImageVersionAlias",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetSagemakerImageVersionArn(val *string) {
	if err := j.validateSetSagemakerImageVersionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sagemakerImageVersionArn",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ResetLifecycleConfigArn() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycleConfigArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ResetSagemakerImageArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSagemakerImageArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ResetSagemakerImageVersionAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetSagemakerImageVersionAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ResetSagemakerImageVersionArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSagemakerImageVersionArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

