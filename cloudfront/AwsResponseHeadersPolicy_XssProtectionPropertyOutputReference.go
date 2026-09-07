package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference interface {
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
	InternalValue() *AwsResponseHeadersPolicy_XssProtectionProperty
	// Experimental.
	SetInternalValue(val *AwsResponseHeadersPolicy_XssProtectionProperty)
	// Experimental.
	ModeBlock() interface{}
	// Experimental.
	SetModeBlock(val interface{})
	// Experimental.
	ModeBlockInput() interface{}
	// Experimental.
	Override() interface{}
	// Experimental.
	SetOverride(val interface{})
	// Experimental.
	OverrideInput() interface{}
	// Experimental.
	Protection() interface{}
	// Experimental.
	SetProtection(val interface{})
	// Experimental.
	ProtectionInput() interface{}
	// Experimental.
	ReportUri() *string
	// Experimental.
	SetReportUri(val *string)
	// Experimental.
	ReportUriInput() *string
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
	ResetModeBlock()
	// Experimental.
	ResetReportUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference
type jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) InternalValue() *AwsResponseHeadersPolicy_XssProtectionProperty {
	var returns *AwsResponseHeadersPolicy_XssProtectionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) ModeBlock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modeBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) ModeBlockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modeBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) OverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) Protection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) ProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) ReportUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) ReportUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsResponseHeadersPolicy_XssProtectionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsResponseHeadersPolicy_XssProtectionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsResponseHeadersPolicy.XssProtectionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsResponseHeadersPolicy_XssProtectionPropertyOutputReference_Override(a AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsResponseHeadersPolicy.XssProtectionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference)SetInternalValue(val *AwsResponseHeadersPolicy_XssProtectionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference)SetModeBlock(val interface{}) {
	if err := j.validateSetModeBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modeBlock",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference)SetOverride(val interface{}) {
	if err := j.validateSetOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"override",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference)SetProtection(val interface{}) {
	if err := j.validateSetProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protection",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference)SetReportUri(val *string) {
	if err := j.validateSetReportUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportUri",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) ResetModeBlock() {
	_jsii_.InvokeVoid(
		a,
		"resetModeBlock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) ResetReportUri() {
	_jsii_.InvokeVoid(
		a,
		"resetReportUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

