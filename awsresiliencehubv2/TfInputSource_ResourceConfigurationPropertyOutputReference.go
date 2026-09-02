package awsresiliencehubv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsresiliencehubv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsresiliencehubv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfInputSource_ResourceConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CfnStackArn() *string
	// Experimental.
	SetCfnStackArn(val *string)
	// Experimental.
	CfnStackArnInput() *string
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
	DesignFileS3Url() *string
	// Experimental.
	SetDesignFileS3Url(val *string)
	// Experimental.
	DesignFileS3UrlInput() *string
	// Experimental.
	Eks() TfInputSource_EksPropertyList
	// Experimental.
	EksInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ResourceTag() TfInputSource_ResourceTagPropertyList
	// Experimental.
	ResourceTagInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TfStateFileUrl() *string
	// Experimental.
	SetTfStateFileUrl(val *string)
	// Experimental.
	TfStateFileUrlInput() *string
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
	PutEks(value interface{})
	// Experimental.
	PutResourceTag(value interface{})
	// Experimental.
	ResetCfnStackArn()
	// Experimental.
	ResetDesignFileS3Url()
	// Experimental.
	ResetEks()
	// Experimental.
	ResetResourceTag()
	// Experimental.
	ResetTfStateFileUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfInputSource_ResourceConfigurationPropertyOutputReference
type jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) CfnStackArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnStackArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) CfnStackArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnStackArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) DesignFileS3Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"designFileS3Url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) DesignFileS3UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"designFileS3UrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) Eks() TfInputSource_EksPropertyList {
	var returns TfInputSource_EksPropertyList
	_jsii_.Get(
		j,
		"eks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) EksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) ResourceTag() TfInputSource_ResourceTagPropertyList {
	var returns TfInputSource_ResourceTagPropertyList
	_jsii_.Get(
		j,
		"resourceTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) ResourceTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) TfStateFileUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tfStateFileUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) TfStateFileUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tfStateFileUrlInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfInputSource_ResourceConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfInputSource_ResourceConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfInputSource_ResourceConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-resilience-hub-v2.TfInputSource.ResourceConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfInputSource_ResourceConfigurationPropertyOutputReference_Override(t TfInputSource_ResourceConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-resilience-hub-v2.TfInputSource.ResourceConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference)SetCfnStackArn(val *string) {
	if err := j.validateSetCfnStackArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cfnStackArn",
		val,
	)
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference)SetDesignFileS3Url(val *string) {
	if err := j.validateSetDesignFileS3UrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"designFileS3Url",
		val,
	)
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference)SetTfStateFileUrl(val *string) {
	if err := j.validateSetTfStateFileUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tfStateFileUrl",
		val,
	)
}

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) PutEks(value interface{}) {
	if err := t.validatePutEksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEks",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) PutResourceTag(value interface{}) {
	if err := t.validatePutResourceTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceTag",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) ResetCfnStackArn() {
	_jsii_.InvokeVoid(
		t,
		"resetCfnStackArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) ResetDesignFileS3Url() {
	_jsii_.InvokeVoid(
		t,
		"resetDesignFileS3Url",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) ResetEks() {
	_jsii_.InvokeVoid(
		t,
		"resetEks",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) ResetResourceTag() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceTag",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) ResetTfStateFileUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetTfStateFileUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfInputSource_ResourceConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

