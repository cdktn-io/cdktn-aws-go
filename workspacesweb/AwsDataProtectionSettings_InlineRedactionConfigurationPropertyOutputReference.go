package workspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/workspacesweb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/workspacesweb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference interface {
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
	GlobalConfidenceLevel() *float64
	// Experimental.
	SetGlobalConfidenceLevel(val *float64)
	// Experimental.
	GlobalConfidenceLevelInput() *float64
	// Experimental.
	GlobalEnforcedUrls() *[]*string
	// Experimental.
	SetGlobalEnforcedUrls(val *[]*string)
	// Experimental.
	GlobalEnforcedUrlsInput() *[]*string
	// Experimental.
	GlobalExemptUrls() *[]*string
	// Experimental.
	SetGlobalExemptUrls(val *[]*string)
	// Experimental.
	GlobalExemptUrlsInput() *[]*string
	// Experimental.
	InlineRedactionPattern() AwsDataProtectionSettings_InlineRedactionPatternPropertyList
	// Experimental.
	InlineRedactionPatternInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
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
	PutInlineRedactionPattern(value interface{})
	// Experimental.
	ResetGlobalConfidenceLevel()
	// Experimental.
	ResetGlobalEnforcedUrls()
	// Experimental.
	ResetGlobalExemptUrls()
	// Experimental.
	ResetInlineRedactionPattern()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference
type jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GlobalConfidenceLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"globalConfidenceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GlobalConfidenceLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"globalConfidenceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GlobalEnforcedUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"globalEnforcedUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GlobalEnforcedUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"globalEnforcedUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GlobalExemptUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"globalExemptUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GlobalExemptUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"globalExemptUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) InlineRedactionPattern() AwsDataProtectionSettings_InlineRedactionPatternPropertyList {
	var returns AwsDataProtectionSettings_InlineRedactionPatternPropertyList
	_jsii_.Get(
		j,
		"inlineRedactionPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) InlineRedactionPatternInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inlineRedactionPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-workspaces-web.AwsDataProtectionSettings.InlineRedactionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference_Override(a AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-workspaces-web.AwsDataProtectionSettings.InlineRedactionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference)SetGlobalConfidenceLevel(val *float64) {
	if err := j.validateSetGlobalConfidenceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalConfidenceLevel",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference)SetGlobalEnforcedUrls(val *[]*string) {
	if err := j.validateSetGlobalEnforcedUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalEnforcedUrls",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference)SetGlobalExemptUrls(val *[]*string) {
	if err := j.validateSetGlobalExemptUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalExemptUrls",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) PutInlineRedactionPattern(value interface{}) {
	if err := a.validatePutInlineRedactionPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInlineRedactionPattern",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) ResetGlobalConfidenceLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetGlobalConfidenceLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) ResetGlobalEnforcedUrls() {
	_jsii_.InvokeVoid(
		a,
		"resetGlobalEnforcedUrls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) ResetGlobalExemptUrls() {
	_jsii_.InvokeVoid(
		a,
		"resetGlobalExemptUrls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) ResetInlineRedactionPattern() {
	_jsii_.InvokeVoid(
		a,
		"resetInlineRedactionPattern",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

