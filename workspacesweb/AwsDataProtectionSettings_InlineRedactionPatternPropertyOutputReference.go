package workspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/workspacesweb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/workspacesweb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BuiltInPatternId() *string
	// Experimental.
	SetBuiltInPatternId(val *string)
	// Experimental.
	BuiltInPatternIdInput() *string
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
	ConfidenceLevel() *float64
	// Experimental.
	SetConfidenceLevel(val *float64)
	// Experimental.
	ConfidenceLevelInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomPattern() AwsDataProtectionSettings_CustomPatternPropertyList
	// Experimental.
	CustomPatternInput() interface{}
	// Experimental.
	EnforcedUrls() *[]*string
	// Experimental.
	SetEnforcedUrls(val *[]*string)
	// Experimental.
	EnforcedUrlsInput() *[]*string
	// Experimental.
	ExemptUrls() *[]*string
	// Experimental.
	SetExemptUrls(val *[]*string)
	// Experimental.
	ExemptUrlsInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RedactionPlaceHolder() AwsDataProtectionSettings_RedactionPlaceHolderPropertyList
	// Experimental.
	RedactionPlaceHolderInput() interface{}
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
	PutCustomPattern(value interface{})
	// Experimental.
	PutRedactionPlaceHolder(value interface{})
	// Experimental.
	ResetBuiltInPatternId()
	// Experimental.
	ResetConfidenceLevel()
	// Experimental.
	ResetCustomPattern()
	// Experimental.
	ResetEnforcedUrls()
	// Experimental.
	ResetExemptUrls()
	// Experimental.
	ResetRedactionPlaceHolder()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference
type jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) BuiltInPatternId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInPatternId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) BuiltInPatternIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInPatternIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) ConfidenceLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confidenceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) ConfidenceLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confidenceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) CustomPattern() AwsDataProtectionSettings_CustomPatternPropertyList {
	var returns AwsDataProtectionSettings_CustomPatternPropertyList
	_jsii_.Get(
		j,
		"customPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) CustomPatternInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) EnforcedUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enforcedUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) EnforcedUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enforcedUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) ExemptUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exemptUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) ExemptUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exemptUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) RedactionPlaceHolder() AwsDataProtectionSettings_RedactionPlaceHolderPropertyList {
	var returns AwsDataProtectionSettings_RedactionPlaceHolderPropertyList
	_jsii_.Get(
		j,
		"redactionPlaceHolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) RedactionPlaceHolderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redactionPlaceHolderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-workspaces-web.AwsDataProtectionSettings.InlineRedactionPatternPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference_Override(a AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-workspaces-web.AwsDataProtectionSettings.InlineRedactionPatternPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference)SetBuiltInPatternId(val *string) {
	if err := j.validateSetBuiltInPatternIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtInPatternId",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference)SetConfidenceLevel(val *float64) {
	if err := j.validateSetConfidenceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidenceLevel",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference)SetEnforcedUrls(val *[]*string) {
	if err := j.validateSetEnforcedUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforcedUrls",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference)SetExemptUrls(val *[]*string) {
	if err := j.validateSetExemptUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exemptUrls",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) PutCustomPattern(value interface{}) {
	if err := a.validatePutCustomPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomPattern",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) PutRedactionPlaceHolder(value interface{}) {
	if err := a.validatePutRedactionPlaceHolderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedactionPlaceHolder",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) ResetBuiltInPatternId() {
	_jsii_.InvokeVoid(
		a,
		"resetBuiltInPatternId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) ResetConfidenceLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetConfidenceLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) ResetCustomPattern() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomPattern",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) ResetEnforcedUrls() {
	_jsii_.InvokeVoid(
		a,
		"resetEnforcedUrls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) ResetExemptUrls() {
	_jsii_.InvokeVoid(
		a,
		"resetExemptUrls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) ResetRedactionPlaceHolder() {
	_jsii_.InvokeVoid(
		a,
		"resetRedactionPlaceHolder",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDataProtectionSettings_InlineRedactionPatternPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

