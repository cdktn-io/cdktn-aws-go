package chimesdkmediapipelines

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/chimesdkmediapipelines/jsii"

	"github.com/cdktn-io/cdktn-aws-go/chimesdkmediapipelines/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference interface {
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
	InternalValue() *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty)
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
	PostCallAnalyticsSettings() AwsMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference
	// Experimental.
	PostCallAnalyticsSettingsInput() *AwsMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty
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
	PutPostCallAnalyticsSettings(value *AwsMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty)
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

// The jsii proxy struct for AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference
type jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) CallAnalyticsStreamCategories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callAnalyticsStreamCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) CallAnalyticsStreamCategoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callAnalyticsStreamCategoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ContentIdentificationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentIdentificationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ContentIdentificationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentIdentificationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ContentRedactionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentRedactionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ContentRedactionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentRedactionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) EnablePartialResultsStabilization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartialResultsStabilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) EnablePartialResultsStabilizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartialResultsStabilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) FilterPartialResults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterPartialResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) FilterPartialResultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterPartialResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) InternalValue() *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty {
	var returns *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) LanguageModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) LanguageModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) PartialResultsStability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partialResultsStability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) PartialResultsStabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partialResultsStabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) PiiEntityTypes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"piiEntityTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) PiiEntityTypesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"piiEntityTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) PostCallAnalyticsSettings() AwsMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference {
	var returns AwsMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"postCallAnalyticsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) PostCallAnalyticsSettingsInput() *AwsMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty {
	var returns *AwsMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty
	_jsii_.Get(
		j,
		"postCallAnalyticsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) VocabularyFilterMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) VocabularyFilterMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) VocabularyFilterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) VocabularyFilterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) VocabularyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) VocabularyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyNameInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference_Override(a AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetCallAnalyticsStreamCategories(val *[]*string) {
	if err := j.validateSetCallAnalyticsStreamCategoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callAnalyticsStreamCategories",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetContentIdentificationType(val *string) {
	if err := j.validateSetContentIdentificationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentIdentificationType",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetContentRedactionType(val *string) {
	if err := j.validateSetContentRedactionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentRedactionType",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetEnablePartialResultsStabilization(val interface{}) {
	if err := j.validateSetEnablePartialResultsStabilizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePartialResultsStabilization",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetFilterPartialResults(val interface{}) {
	if err := j.validateSetFilterPartialResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterPartialResults",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetInternalValue(val *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetLanguageModelName(val *string) {
	if err := j.validateSetLanguageModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageModelName",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetPartialResultsStability(val *string) {
	if err := j.validateSetPartialResultsStabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partialResultsStability",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetPiiEntityTypes(val *string) {
	if err := j.validateSetPiiEntityTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"piiEntityTypes",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetVocabularyFilterMethod(val *string) {
	if err := j.validateSetVocabularyFilterMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyFilterMethod",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetVocabularyFilterName(val *string) {
	if err := j.validateSetVocabularyFilterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyFilterName",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)SetVocabularyName(val *string) {
	if err := j.validateSetVocabularyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyName",
		val,
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) PutPostCallAnalyticsSettings(value *AwsMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty) {
	if err := a.validatePutPostCallAnalyticsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPostCallAnalyticsSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetCallAnalyticsStreamCategories() {
	_jsii_.InvokeVoid(
		a,
		"resetCallAnalyticsStreamCategories",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetContentIdentificationType() {
	_jsii_.InvokeVoid(
		a,
		"resetContentIdentificationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetContentRedactionType() {
	_jsii_.InvokeVoid(
		a,
		"resetContentRedactionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetEnablePartialResultsStabilization() {
	_jsii_.InvokeVoid(
		a,
		"resetEnablePartialResultsStabilization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetFilterPartialResults() {
	_jsii_.InvokeVoid(
		a,
		"resetFilterPartialResults",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetLanguageModelName() {
	_jsii_.InvokeVoid(
		a,
		"resetLanguageModelName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetPartialResultsStability() {
	_jsii_.InvokeVoid(
		a,
		"resetPartialResultsStability",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetPiiEntityTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetPiiEntityTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetPostCallAnalyticsSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetPostCallAnalyticsSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetVocabularyFilterMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetVocabularyFilterMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetVocabularyFilterName() {
	_jsii_.InvokeVoid(
		a,
		"resetVocabularyFilterName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ResetVocabularyName() {
	_jsii_.InvokeVoid(
		a,
		"resetVocabularyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

