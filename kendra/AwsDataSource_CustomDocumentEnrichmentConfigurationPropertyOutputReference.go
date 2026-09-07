package kendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kendra/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kendra/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference interface {
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
	InlineConfigurations() AwsDataSource_InlineConfigurationsPropertyList
	// Experimental.
	InlineConfigurationsInput() interface{}
	// Experimental.
	InternalValue() *AwsDataSource_CustomDocumentEnrichmentConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsDataSource_CustomDocumentEnrichmentConfigurationProperty)
	// Experimental.
	PostExtractionHookConfiguration() AwsDataSource_PostExtractionHookConfigurationPropertyOutputReference
	// Experimental.
	PostExtractionHookConfigurationInput() *AwsDataSource_PostExtractionHookConfigurationProperty
	// Experimental.
	PreExtractionHookConfiguration() AwsDataSource_PreExtractionHookConfigurationPropertyOutputReference
	// Experimental.
	PreExtractionHookConfigurationInput() *AwsDataSource_PreExtractionHookConfigurationProperty
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
	PutPostExtractionHookConfiguration(value *AwsDataSource_PostExtractionHookConfigurationProperty)
	// Experimental.
	PutPreExtractionHookConfiguration(value *AwsDataSource_PreExtractionHookConfigurationProperty)
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

// The jsii proxy struct for AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference
type jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) InlineConfigurations() AwsDataSource_InlineConfigurationsPropertyList {
	var returns AwsDataSource_InlineConfigurationsPropertyList
	_jsii_.Get(
		j,
		"inlineConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) InlineConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inlineConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) InternalValue() *AwsDataSource_CustomDocumentEnrichmentConfigurationProperty {
	var returns *AwsDataSource_CustomDocumentEnrichmentConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) PostExtractionHookConfiguration() AwsDataSource_PostExtractionHookConfigurationPropertyOutputReference {
	var returns AwsDataSource_PostExtractionHookConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"postExtractionHookConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) PostExtractionHookConfigurationInput() *AwsDataSource_PostExtractionHookConfigurationProperty {
	var returns *AwsDataSource_PostExtractionHookConfigurationProperty
	_jsii_.Get(
		j,
		"postExtractionHookConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) PreExtractionHookConfiguration() AwsDataSource_PreExtractionHookConfigurationPropertyOutputReference {
	var returns AwsDataSource_PreExtractionHookConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"preExtractionHookConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) PreExtractionHookConfigurationInput() *AwsDataSource_PreExtractionHookConfigurationProperty {
	var returns *AwsDataSource_PreExtractionHookConfigurationProperty
	_jsii_.Get(
		j,
		"preExtractionHookConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kendra.AwsDataSource.CustomDocumentEnrichmentConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference_Override(a AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kendra.AwsDataSource.CustomDocumentEnrichmentConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference)SetInternalValue(val *AwsDataSource_CustomDocumentEnrichmentConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) PutInlineConfigurations(value interface{}) {
	if err := a.validatePutInlineConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInlineConfigurations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) PutPostExtractionHookConfiguration(value *AwsDataSource_PostExtractionHookConfigurationProperty) {
	if err := a.validatePutPostExtractionHookConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPostExtractionHookConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) PutPreExtractionHookConfiguration(value *AwsDataSource_PreExtractionHookConfigurationProperty) {
	if err := a.validatePutPreExtractionHookConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPreExtractionHookConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ResetInlineConfigurations() {
	_jsii_.InvokeVoid(
		a,
		"resetInlineConfigurations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ResetPostExtractionHookConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetPostExtractionHookConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ResetPreExtractionHookConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetPreExtractionHookConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDataSource_CustomDocumentEnrichmentConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

