package chimesdkmediapipelines

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/chimesdkmediapipelines/jsii"

	"github.com/cdktn-io/cdktn-aws-go/chimesdkmediapipelines/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference interface {
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
	InternalValue() *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty)
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
	ShowSpeakerLabel() interface{}
	// Experimental.
	SetShowSpeakerLabel(val interface{})
	// Experimental.
	ShowSpeakerLabelInput() interface{}
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
	ResetShowSpeakerLabel()
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

// The jsii proxy struct for AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference
type jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ContentIdentificationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentIdentificationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ContentIdentificationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentIdentificationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ContentRedactionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentRedactionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ContentRedactionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentRedactionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) EnablePartialResultsStabilization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartialResultsStabilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) EnablePartialResultsStabilizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartialResultsStabilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) FilterPartialResults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterPartialResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) FilterPartialResultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterPartialResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) InternalValue() *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty {
	var returns *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) LanguageModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) LanguageModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) PartialResultsStability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partialResultsStability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) PartialResultsStabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partialResultsStabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) PiiEntityTypes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"piiEntityTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) PiiEntityTypesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"piiEntityTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ShowSpeakerLabel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showSpeakerLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ShowSpeakerLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showSpeakerLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) VocabularyFilterMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) VocabularyFilterMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) VocabularyFilterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) VocabularyFilterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) VocabularyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) VocabularyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyNameInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.AmazonTranscribeProcessorConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference_Override(a AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.AmazonTranscribeProcessorConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetContentIdentificationType(val *string) {
	if err := j.validateSetContentIdentificationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentIdentificationType",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetContentRedactionType(val *string) {
	if err := j.validateSetContentRedactionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentRedactionType",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetEnablePartialResultsStabilization(val interface{}) {
	if err := j.validateSetEnablePartialResultsStabilizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePartialResultsStabilization",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetFilterPartialResults(val interface{}) {
	if err := j.validateSetFilterPartialResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterPartialResults",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetInternalValue(val *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetLanguageModelName(val *string) {
	if err := j.validateSetLanguageModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageModelName",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetPartialResultsStability(val *string) {
	if err := j.validateSetPartialResultsStabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partialResultsStability",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetPiiEntityTypes(val *string) {
	if err := j.validateSetPiiEntityTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"piiEntityTypes",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetShowSpeakerLabel(val interface{}) {
	if err := j.validateSetShowSpeakerLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showSpeakerLabel",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetVocabularyFilterMethod(val *string) {
	if err := j.validateSetVocabularyFilterMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyFilterMethod",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetVocabularyFilterName(val *string) {
	if err := j.validateSetVocabularyFilterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyFilterName",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)SetVocabularyName(val *string) {
	if err := j.validateSetVocabularyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyName",
		val,
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ResetContentIdentificationType() {
	_jsii_.InvokeVoid(
		a,
		"resetContentIdentificationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ResetContentRedactionType() {
	_jsii_.InvokeVoid(
		a,
		"resetContentRedactionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ResetEnablePartialResultsStabilization() {
	_jsii_.InvokeVoid(
		a,
		"resetEnablePartialResultsStabilization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ResetFilterPartialResults() {
	_jsii_.InvokeVoid(
		a,
		"resetFilterPartialResults",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ResetLanguageModelName() {
	_jsii_.InvokeVoid(
		a,
		"resetLanguageModelName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ResetPartialResultsStability() {
	_jsii_.InvokeVoid(
		a,
		"resetPartialResultsStability",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ResetPiiEntityTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetPiiEntityTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ResetShowSpeakerLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetShowSpeakerLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ResetVocabularyFilterMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetVocabularyFilterMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ResetVocabularyFilterName() {
	_jsii_.InvokeVoid(
		a,
		"resetVocabularyFilterName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ResetVocabularyName() {
	_jsii_.InvokeVoid(
		a,
		"resetVocabularyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

