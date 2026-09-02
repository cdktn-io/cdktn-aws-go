package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessControlMaxAgeSec() *float64
	// Experimental.
	SetAccessControlMaxAgeSec(val *float64)
	// Experimental.
	AccessControlMaxAgeSecInput() *float64
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
	IncludeSubdomains() interface{}
	// Experimental.
	SetIncludeSubdomains(val interface{})
	// Experimental.
	IncludeSubdomainsInput() interface{}
	// Experimental.
	InternalValue() *TfResponseHeadersPolicy_StrictTransportSecurityProperty
	// Experimental.
	SetInternalValue(val *TfResponseHeadersPolicy_StrictTransportSecurityProperty)
	// Experimental.
	Override() interface{}
	// Experimental.
	SetOverride(val interface{})
	// Experimental.
	OverrideInput() interface{}
	// Experimental.
	Preload() interface{}
	// Experimental.
	SetPreload(val interface{})
	// Experimental.
	PreloadInput() interface{}
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
	ResetIncludeSubdomains()
	// Experimental.
	ResetPreload()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference
type jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) AccessControlMaxAgeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) AccessControlMaxAgeSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) IncludeSubdomains() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubdomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) IncludeSubdomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubdomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) InternalValue() *TfResponseHeadersPolicy_StrictTransportSecurityProperty {
	var returns *TfResponseHeadersPolicy_StrictTransportSecurityProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) OverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) Preload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) PreloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfResponseHeadersPolicy.StrictTransportSecurityPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference_Override(t TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfResponseHeadersPolicy.StrictTransportSecurityPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference)SetAccessControlMaxAgeSec(val *float64) {
	if err := j.validateSetAccessControlMaxAgeSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessControlMaxAgeSec",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference)SetIncludeSubdomains(val interface{}) {
	if err := j.validateSetIncludeSubdomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSubdomains",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference)SetInternalValue(val *TfResponseHeadersPolicy_StrictTransportSecurityProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference)SetOverride(val interface{}) {
	if err := j.validateSetOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"override",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference)SetPreload(val interface{}) {
	if err := j.validateSetPreloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preload",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) ResetIncludeSubdomains() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeSubdomains",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) ResetPreload() {
	_jsii_.InvokeVoid(
		t,
		"resetPreload",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

