package awsresiliencehubv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsresiliencehubv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsresiliencehubv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference interface {
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
	Eks() AwsResiliencehubv2InputSource_EksPropertyList
	// Experimental.
	EksInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ResourceTag() AwsResiliencehubv2InputSource_ResourceTagPropertyList
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

// The jsii proxy struct for AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference
type jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) CfnStackArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnStackArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) CfnStackArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnStackArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) DesignFileS3Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"designFileS3Url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) DesignFileS3UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"designFileS3UrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) Eks() AwsResiliencehubv2InputSource_EksPropertyList {
	var returns AwsResiliencehubv2InputSource_EksPropertyList
	_jsii_.Get(
		j,
		"eks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) EksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) ResourceTag() AwsResiliencehubv2InputSource_ResourceTagPropertyList {
	var returns AwsResiliencehubv2InputSource_ResourceTagPropertyList
	_jsii_.Get(
		j,
		"resourceTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) ResourceTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) TfStateFileUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tfStateFileUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) TfStateFileUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tfStateFileUrlInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-resilience-hub-v2.AwsResiliencehubv2InputSource.ResourceConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference_Override(a AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-resilience-hub-v2.AwsResiliencehubv2InputSource.ResourceConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference)SetCfnStackArn(val *string) {
	if err := j.validateSetCfnStackArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cfnStackArn",
		val,
	)
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference)SetDesignFileS3Url(val *string) {
	if err := j.validateSetDesignFileS3UrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"designFileS3Url",
		val,
	)
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference)SetTfStateFileUrl(val *string) {
	if err := j.validateSetTfStateFileUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tfStateFileUrl",
		val,
	)
}

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) PutEks(value interface{}) {
	if err := a.validatePutEksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEks",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) PutResourceTag(value interface{}) {
	if err := a.validatePutResourceTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceTag",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) ResetCfnStackArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCfnStackArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) ResetDesignFileS3Url() {
	_jsii_.InvokeVoid(
		a,
		"resetDesignFileS3Url",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) ResetEks() {
	_jsii_.InvokeVoid(
		a,
		"resetEks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) ResetResourceTag() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceTag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) ResetTfStateFileUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetTfStateFileUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsResiliencehubv2InputSource_ResourceConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

