package awschimesdkmediapipelines

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awschimesdkmediapipelines/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awschimesdkmediapipelines/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CallAnalyticsStreamCategories() *[]*string
	// Experimental.
	SetCallAnalyticsStreamCategories(val *[]*string)
	// Experimental.
	CallAnalyticsStreamCategoriesInput() *[]*string
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
	ContentIdentificationType() *string
	// Experimental.
	SetContentIdentificationType(val *string)
	// Experimental.
	ContentIdentificationTypeInput() *string
	// Experimental.
	ContentRedactionType() *string
	// Experimental.
	SetContentRedactionType(val *string)
	// Experimental.
	ContentRedactionTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EnablePartialResultsStabilization() interface{}
	// Experimental.
	SetEnablePartialResultsStabilization(val interface{})
	// Experimental.
	EnablePartialResultsStabilizationInput() interface{}
	// Experimental.
	FilterPartialResults() interface{}
	// Experimental.
	SetFilterPartialResults(val interface{})
	// Experimental.
	FilterPartialResultsInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty)
	// Experimental.
	LanguageCode() *string
	// Experimental.
	SetLanguageCode(val *string)
	// Experimental.
	LanguageCodeInput() *string
	// Experimental.
	LanguageModelName() *string
	// Experimental.
	SetLanguageModelName(val *string)
	// Experimental.
	LanguageModelNameInput() *string
	// Experimental.
	PartialResultsStability() *string
	// Experimental.
	SetPartialResultsStability(val *string)
	// Experimental.
	PartialResultsStabilityInput() *string
	// Experimental.
	PiiEntityTypes() *string
	// Experimental.
	SetPiiEntityTypes(val *string)
	// Experimental.
	PiiEntityTypesInput() *string
	// Experimental.
	PostCallAnalyticsSettings() TfMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference
	// Experimental.
	PostCallAnalyticsSettingsInput() *TfMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VocabularyFilterMethod() *string
	// Experimental.
	SetVocabularyFilterMethod(val *string)
	// Experimental.
	VocabularyFilterMethodInput() *string
	// Experimental.
	VocabularyFilterName() *string
	// Experimental.
	SetVocabularyFilterName(val *string)
	// Experimental.
	VocabularyFilterNameInput() *string
	// Experimental.
	VocabularyName() *string
	// Experimental.
	SetVocabularyName(val *string)
	// Experimental.
	VocabularyNameInput() *string
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
	PutPostCallAnalyticsSettings(value *TfMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty)
	// Experimental.
	ResetCallAnalyticsStreamCategories()
	// Experimental.
	ResetContentIdentificationType()
	// Experimental.
	ResetContentRedactionType()
	// Experimental.
	ResetEnablePartialResultsStabilization()
	// Experimental.
	ResetFilterPartialResults()
	// Experimental.
	ResetLanguageModelName()
	// Experimental.
	ResetPartialResultsStability()
	// Experimental.
	ResetPiiEntityTypes()
	// Experimental.
	ResetPostCallAnalyticsSettings()
	// Experimental.
	ResetVocabularyFilterMethod()
	// Experimental.
	ResetVocabularyFilterName()
	// Experimental.
	ResetVocabularyName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference
type jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) CallAnalyticsStreamCategories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callAnalyticsStreamCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) CallAnalyticsStreamCategoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callAnalyticsStreamCategoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ContentIdentificationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentIdentificationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ContentIdentificationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentIdentificationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ContentRedactionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentRedactionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ContentRedactionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentRedactionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) EnablePartialResultsStabilization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartialResultsStabilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) EnablePartialResultsStabilizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartialResultsStabilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) FilterPartialResults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterPartialResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) FilterPartialResultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterPartialResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) InternalValue() *TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty {
	var returns *TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) LanguageModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) LanguageModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) PartialResultsStability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partialResultsStability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) PartialResultsStabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partialResultsStabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) PiiEntityTypes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"piiEntityTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) PiiEntityTypesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"piiEntityTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) PostCallAnalyticsSettings() TfMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference {
	var returns TfMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"postCallAnalyticsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) PostCallAnalyticsSettingsInput() *TfMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty {
	var returns *TfMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty
	_jsii_.Get(
		j,
		"postCallAnalyticsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) VocabularyFilterMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) VocabularyFilterMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) VocabularyFilterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) VocabularyFilterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) VocabularyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) VocabularyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyNameInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.TfMediaInsightsPipelineConfiguration.AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference_Override(t TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.TfMediaInsightsPipelineConfiguration.AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetCallAnalyticsStreamCategories(val *[]*string) {
	if err := j.validateSetCallAnalyticsStreamCategoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callAnalyticsStreamCategories",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetContentIdentificationType(val *string) {
	if err := j.validateSetContentIdentificationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentIdentificationType",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetContentRedactionType(val *string) {
	if err := j.validateSetContentRedactionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentRedactionType",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetEnablePartialResultsStabilization(val interface{}) {
	if err := j.validateSetEnablePartialResultsStabilizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePartialResultsStabilization",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetFilterPartialResults(val interface{}) {
	if err := j.validateSetFilterPartialResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterPartialResults",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetInternalValue(val *TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetLanguageModelName(val *string) {
	if err := j.validateSetLanguageModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageModelName",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetPartialResultsStability(val *string) {
	if err := j.validateSetPartialResultsStabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partialResultsStability",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetPiiEntityTypes(val *string) {
	if err := j.validateSetPiiEntityTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"piiEntityTypes",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetVocabularyFilterMethod(val *string) {
	if err := j.validateSetVocabularyFilterMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyFilterMethod",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetVocabularyFilterName(val *string) {
	if err := j.validateSetVocabularyFilterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyFilterName",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetVocabularyName(val *string) {
	if err := j.validateSetVocabularyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyName",
		val,
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) PutPostCallAnalyticsSettings(value *TfMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty) {
	if err := t.validatePutPostCallAnalyticsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPostCallAnalyticsSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetCallAnalyticsStreamCategories() {
	_jsii_.InvokeVoid(
		t,
		"resetCallAnalyticsStreamCategories",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetContentIdentificationType() {
	_jsii_.InvokeVoid(
		t,
		"resetContentIdentificationType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetContentRedactionType() {
	_jsii_.InvokeVoid(
		t,
		"resetContentRedactionType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetEnablePartialResultsStabilization() {
	_jsii_.InvokeVoid(
		t,
		"resetEnablePartialResultsStabilization",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetFilterPartialResults() {
	_jsii_.InvokeVoid(
		t,
		"resetFilterPartialResults",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetLanguageModelName() {
	_jsii_.InvokeVoid(
		t,
		"resetLanguageModelName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetPartialResultsStability() {
	_jsii_.InvokeVoid(
		t,
		"resetPartialResultsStability",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetPiiEntityTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetPiiEntityTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetPostCallAnalyticsSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetPostCallAnalyticsSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetVocabularyFilterMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetVocabularyFilterMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetVocabularyFilterName() {
	_jsii_.InvokeVoid(
		t,
		"resetVocabularyFilterName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetVocabularyName() {
	_jsii_.InvokeVoid(
		t,
		"resetVocabularyName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

