package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference interface {
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
	// Experimental.
	CookiesConfig() DataAwsCloudfrontCachePolicy_CookiesConfigPropertyList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EnableAcceptEncodingBrotli() cdktn.IResolvable
	// Experimental.
	EnableAcceptEncodingGzip() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	// Experimental.
	HeadersConfig() DataAwsCloudfrontCachePolicy_HeadersConfigPropertyList
	// Experimental.
	InternalValue() *DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty
	// Experimental.
	SetInternalValue(val *DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty)
	// Experimental.
	QueryStringsConfig() DataAwsCloudfrontCachePolicy_QueryStringsConfigPropertyList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference
type jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) CookiesConfig() DataAwsCloudfrontCachePolicy_CookiesConfigPropertyList {
	var returns DataAwsCloudfrontCachePolicy_CookiesConfigPropertyList
	_jsii_.Get(
		j,
		"cookiesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) EnableAcceptEncodingBrotli() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enableAcceptEncodingBrotli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) EnableAcceptEncodingGzip() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enableAcceptEncodingGzip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) HeadersConfig() DataAwsCloudfrontCachePolicy_HeadersConfigPropertyList {
	var returns DataAwsCloudfrontCachePolicy_HeadersConfigPropertyList
	_jsii_.Get(
		j,
		"headersConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) InternalValue() *DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty {
	var returns *DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) QueryStringsConfig() DataAwsCloudfrontCachePolicy_QueryStringsConfigPropertyList {
	var returns DataAwsCloudfrontCachePolicy_QueryStringsConfigPropertyList
	_jsii_.Get(
		j,
		"queryStringsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.DataAwsCloudfrontCachePolicy.ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference_Override(d DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.DataAwsCloudfrontCachePolicy.ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetInternalValue(val *DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

