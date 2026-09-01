package awskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskendra/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskendra/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference interface {
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
	InlineConfigurations() AwsKendraDataSource_InlineConfigurationsPropertyList
	// Experimental.
	InlineConfigurationsInput() interface{}
	// Experimental.
	InternalValue() *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationProperty)
	// Experimental.
	PostExtractionHookConfiguration() AwsKendraDataSource_PostExtractionHookConfigurationPropertyOutputReference
	// Experimental.
	PostExtractionHookConfigurationInput() *AwsKendraDataSource_PostExtractionHookConfigurationProperty
	// Experimental.
	PreExtractionHookConfiguration() AwsKendraDataSource_PreExtractionHookConfigurationPropertyOutputReference
	// Experimental.
	PreExtractionHookConfigurationInput() *AwsKendraDataSource_PreExtractionHookConfigurationProperty
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
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
	PutInlineConfigurations(value interface{})
	// Experimental.
	PutPostExtractionHookConfiguration(value *AwsKendraDataSource_PostExtractionHookConfigurationProperty)
	// Experimental.
	PutPreExtractionHookConfiguration(value *AwsKendraDataSource_PreExtractionHookConfigurationProperty)
	// Experimental.
	ResetInlineConfigurations()
	// Experimental.
	ResetPostExtractionHookConfiguration()
	// Experimental.
	ResetPreExtractionHookConfiguration()
	// Experimental.
	ResetRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference
type jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) InlineConfigurations() AwsKendraDataSource_InlineConfigurationsPropertyList {
	var returns AwsKendraDataSource_InlineConfigurationsPropertyList
	_jsii_.Get(
		j,
		"inlineConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) InlineConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inlineConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) InternalValue() *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationProperty {
	var returns *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) PostExtractionHookConfiguration() AwsKendraDataSource_PostExtractionHookConfigurationPropertyOutputReference {
	var returns AwsKendraDataSource_PostExtractionHookConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"postExtractionHookConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) PostExtractionHookConfigurationInput() *AwsKendraDataSource_PostExtractionHookConfigurationProperty {
	var returns *AwsKendraDataSource_PostExtractionHookConfigurationProperty
	_jsii_.Get(
		j,
		"postExtractionHookConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) PreExtractionHookConfiguration() AwsKendraDataSource_PreExtractionHookConfigurationPropertyOutputReference {
	var returns AwsKendraDataSource_PreExtractionHookConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"preExtractionHookConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) PreExtractionHookConfigurationInput() *AwsKendraDataSource_PreExtractionHookConfigurationProperty {
	var returns *AwsKendraDataSource_PreExtractionHookConfigurationProperty
	_jsii_.Get(
		j,
		"preExtractionHookConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kendra.AwsKendraDataSource.CustomDocumentEnrichmentConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference_Override(a AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kendra.AwsKendraDataSource.CustomDocumentEnrichmentConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference)SetInternalValue(val *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) PutInlineConfigurations(value interface{}) {
	if err := a.validatePutInlineConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInlineConfigurations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) PutPostExtractionHookConfiguration(value *AwsKendraDataSource_PostExtractionHookConfigurationProperty) {
	if err := a.validatePutPostExtractionHookConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPostExtractionHookConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) PutPreExtractionHookConfiguration(value *AwsKendraDataSource_PreExtractionHookConfigurationProperty) {
	if err := a.validatePutPreExtractionHookConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPreExtractionHookConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ResetInlineConfigurations() {
	_jsii_.InvokeVoid(
		a,
		"resetInlineConfigurations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ResetPostExtractionHookConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetPostExtractionHookConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ResetPreExtractionHookConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetPreExtractionHookConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

