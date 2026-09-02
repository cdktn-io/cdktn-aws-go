package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDistribution_VpcOriginConfigPropertyOutputReference interface {
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
	InternalValue() *TfDistribution_VpcOriginConfigProperty
	// Experimental.
	SetInternalValue(val *TfDistribution_VpcOriginConfigProperty)
	// Experimental.
	OriginKeepaliveTimeout() *float64
	// Experimental.
	SetOriginKeepaliveTimeout(val *float64)
	// Experimental.
	OriginKeepaliveTimeoutInput() *float64
	// Experimental.
	OriginReadTimeout() *float64
	// Experimental.
	SetOriginReadTimeout(val *float64)
	// Experimental.
	OriginReadTimeoutInput() *float64
	// Experimental.
	OwnerAccountId() *string
	// Experimental.
	SetOwnerAccountId(val *string)
	// Experimental.
	OwnerAccountIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VpcOriginId() *string
	// Experimental.
	SetVpcOriginId(val *string)
	// Experimental.
	VpcOriginIdInput() *string
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
	ResetOriginKeepaliveTimeout()
	// Experimental.
	ResetOriginReadTimeout()
	// Experimental.
	ResetOwnerAccountId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDistribution_VpcOriginConfigPropertyOutputReference
type jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) InternalValue() *TfDistribution_VpcOriginConfigProperty {
	var returns *TfDistribution_VpcOriginConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) OriginKeepaliveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originKeepaliveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) OriginKeepaliveTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originKeepaliveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) OriginReadTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originReadTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) OriginReadTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originReadTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) OwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) OwnerAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) VpcOriginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcOriginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) VpcOriginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcOriginIdInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDistribution_VpcOriginConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDistribution_VpcOriginConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDistribution_VpcOriginConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfDistribution.VpcOriginConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDistribution_VpcOriginConfigPropertyOutputReference_Override(t TfDistribution_VpcOriginConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfDistribution.VpcOriginConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference)SetInternalValue(val *TfDistribution_VpcOriginConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference)SetOriginKeepaliveTimeout(val *float64) {
	if err := j.validateSetOriginKeepaliveTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originKeepaliveTimeout",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference)SetOriginReadTimeout(val *float64) {
	if err := j.validateSetOriginReadTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originReadTimeout",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference)SetOwnerAccountId(val *string) {
	if err := j.validateSetOwnerAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownerAccountId",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference)SetVpcOriginId(val *string) {
	if err := j.validateSetVpcOriginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcOriginId",
		val,
	)
}

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) ResetOriginKeepaliveTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetOriginKeepaliveTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) ResetOriginReadTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetOriginReadTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) ResetOwnerAccountId() {
	_jsii_.InvokeVoid(
		t,
		"resetOwnerAccountId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDistribution_VpcOriginConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

