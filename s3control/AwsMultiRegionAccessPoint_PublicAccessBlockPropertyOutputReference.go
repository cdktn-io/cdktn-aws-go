package s3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/s3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference interface {
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
	InternalValue() *AwsMultiRegionAccessPoint_PublicAccessBlockProperty
	// Experimental.
	SetInternalValue(val *AwsMultiRegionAccessPoint_PublicAccessBlockProperty)
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

// The jsii proxy struct for AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference
type jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) BlockPublicAcls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) BlockPublicAclsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicAclsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) BlockPublicPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) BlockPublicPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) IgnorePublicAcls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignorePublicAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) IgnorePublicAclsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignorePublicAclsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) InternalValue() *AwsMultiRegionAccessPoint_PublicAccessBlockProperty {
	var returns *AwsMultiRegionAccessPoint_PublicAccessBlockProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) RestrictPublicBuckets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictPublicBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) RestrictPublicBucketsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictPublicBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsMultiRegionAccessPoint.PublicAccessBlockPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference_Override(a AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsMultiRegionAccessPoint.PublicAccessBlockPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference)SetBlockPublicAcls(val interface{}) {
	if err := j.validateSetBlockPublicAclsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockPublicAcls",
		val,
	)
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference)SetBlockPublicPolicy(val interface{}) {
	if err := j.validateSetBlockPublicPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockPublicPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference)SetIgnorePublicAcls(val interface{}) {
	if err := j.validateSetIgnorePublicAclsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignorePublicAcls",
		val,
	)
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference)SetInternalValue(val *AwsMultiRegionAccessPoint_PublicAccessBlockProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference)SetRestrictPublicBuckets(val interface{}) {
	if err := j.validateSetRestrictPublicBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictPublicBuckets",
		val,
	)
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) ResetBlockPublicAcls() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockPublicAcls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) ResetBlockPublicPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockPublicPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) ResetIgnorePublicAcls() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnorePublicAcls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) ResetRestrictPublicBuckets() {
	_jsii_.InvokeVoid(
		a,
		"resetRestrictPublicBuckets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMultiRegionAccessPoint_PublicAccessBlockPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

