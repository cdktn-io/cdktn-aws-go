package awschimesdkmediapipelines

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awschimesdkmediapipelines/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awschimesdkmediapipelines/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference interface {
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
	InternalValue() *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty)
	// Experimental.
	SpeakerSearchStatus() *string
	// Experimental.
	SetSpeakerSearchStatus(val *string)
	// Experimental.
	SpeakerSearchStatusInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VoiceToneAnalysisStatus() *string
	// Experimental.
	SetVoiceToneAnalysisStatus(val *string)
	// Experimental.
	VoiceToneAnalysisStatusInput() *string
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

// The jsii proxy struct for AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference
type jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) InternalValue() *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty {
	var returns *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) SpeakerSearchStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"speakerSearchStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) SpeakerSearchStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"speakerSearchStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) VoiceToneAnalysisStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceToneAnalysisStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) VoiceToneAnalysisStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceToneAnalysisStatusInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration.VoiceAnalyticsProcessorConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference_Override(a AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration.VoiceAnalyticsProcessorConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference)SetInternalValue(val *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference)SetSpeakerSearchStatus(val *string) {
	if err := j.validateSetSpeakerSearchStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"speakerSearchStatus",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference)SetVoiceToneAnalysisStatus(val *string) {
	if err := j.validateSetVoiceToneAnalysisStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"voiceToneAnalysisStatus",
		val,
	)
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

