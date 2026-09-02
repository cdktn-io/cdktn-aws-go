package awss3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BlockPublicAcls() interface{}
	// Experimental.
	SetBlockPublicAcls(val interface{})
	// Experimental.
	BlockPublicAclsInput() interface{}
	// Experimental.
	BlockPublicPolicy() interface{}
	// Experimental.
	SetBlockPublicPolicy(val interface{})
	// Experimental.
	BlockPublicPolicyInput() interface{}
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
	IgnorePublicAcls() interface{}
	// Experimental.
	SetIgnorePublicAcls(val interface{})
	// Experimental.
	IgnorePublicAclsInput() interface{}
	// Experimental.
	InternalValue() *TfAccessPoint_PublicAccessBlockConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfAccessPoint_PublicAccessBlockConfigurationProperty)
	// Experimental.
	RestrictPublicBuckets() interface{}
	// Experimental.
	SetRestrictPublicBuckets(val interface{})
	// Experimental.
	RestrictPublicBucketsInput() interface{}
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
	ResetBlockPublicAcls()
	// Experimental.
	ResetBlockPublicPolicy()
	// Experimental.
	ResetIgnorePublicAcls()
	// Experimental.
	ResetRestrictPublicBuckets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference
type jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) BlockPublicAcls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) BlockPublicAclsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicAclsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) BlockPublicPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) BlockPublicPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) IgnorePublicAcls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignorePublicAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) IgnorePublicAclsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignorePublicAclsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) InternalValue() *TfAccessPoint_PublicAccessBlockConfigurationProperty {
	var returns *TfAccessPoint_PublicAccessBlockConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) RestrictPublicBuckets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictPublicBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) RestrictPublicBucketsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictPublicBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.TfAccessPoint.PublicAccessBlockConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference_Override(t TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.TfAccessPoint.PublicAccessBlockConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference)SetBlockPublicAcls(val interface{}) {
	if err := j.validateSetBlockPublicAclsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockPublicAcls",
		val,
	)
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference)SetBlockPublicPolicy(val interface{}) {
	if err := j.validateSetBlockPublicPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockPublicPolicy",
		val,
	)
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference)SetIgnorePublicAcls(val interface{}) {
	if err := j.validateSetIgnorePublicAclsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignorePublicAcls",
		val,
	)
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference)SetInternalValue(val *TfAccessPoint_PublicAccessBlockConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference)SetRestrictPublicBuckets(val interface{}) {
	if err := j.validateSetRestrictPublicBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictPublicBuckets",
		val,
	)
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) ResetBlockPublicAcls() {
	_jsii_.InvokeVoid(
		t,
		"resetBlockPublicAcls",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) ResetBlockPublicPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetBlockPublicPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) ResetIgnorePublicAcls() {
	_jsii_.InvokeVoid(
		t,
		"resetIgnorePublicAcls",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) ResetRestrictPublicBuckets() {
	_jsii_.InvokeVoid(
		t,
		"resetRestrictPublicBuckets",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAccessPoint_PublicAccessBlockConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

