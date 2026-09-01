package awschimesdkmediapipelines

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awschimesdkmediapipelines/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awschimesdkmediapipelines/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IssueDetectionConfiguration() AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_IssueDetectionConfigurationPropertyOutputReference
	// Experimental.
	IssueDetectionConfigurationInput() *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty
	// Experimental.
	KeywordMatchConfiguration() AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_KeywordMatchConfigurationPropertyOutputReference
	// Experimental.
	KeywordMatchConfigurationInput() *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty
	// Experimental.
	SentimentConfiguration() AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference
	// Experimental.
	SentimentConfigurationInput() *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	PutIssueDetectionConfiguration(value *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty)
	// Experimental.
	PutKeywordMatchConfiguration(value *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty)
	// Experimental.
	PutSentimentConfiguration(value *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationProperty)
	// Experimental.
	ResetIssueDetectionConfiguration()
	// Experimental.
	ResetKeywordMatchConfiguration()
	// Experimental.
	ResetSentimentConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference
type jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) IssueDetectionConfiguration() AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_IssueDetectionConfigurationPropertyOutputReference {
	var returns AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_IssueDetectionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"issueDetectionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) IssueDetectionConfigurationInput() *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty {
	var returns *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty
	_jsii_.Get(
		j,
		"issueDetectionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) KeywordMatchConfiguration() AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_KeywordMatchConfigurationPropertyOutputReference {
	var returns AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_KeywordMatchConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"keywordMatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) KeywordMatchConfigurationInput() *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty {
	var returns *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty
	_jsii_.Get(
		j,
		"keywordMatchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) SentimentConfiguration() AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference {
	var returns AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"sentimentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) SentimentConfigurationInput() *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationProperty {
	var returns *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationProperty
	_jsii_.Get(
		j,
		"sentimentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration.RulesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference_Override(a AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration.RulesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) PutIssueDetectionConfiguration(value *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty) {
	if err := a.validatePutIssueDetectionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIssueDetectionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) PutKeywordMatchConfiguration(value *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty) {
	if err := a.validatePutKeywordMatchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKeywordMatchConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) PutSentimentConfiguration(value *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationProperty) {
	if err := a.validatePutSentimentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSentimentConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ResetIssueDetectionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetIssueDetectionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ResetKeywordMatchConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetKeywordMatchConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ResetSentimentConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSentimentConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

