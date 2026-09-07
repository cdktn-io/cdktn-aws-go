package chimesdkmediapipelines

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/chimesdkmediapipelines/jsii"

	"github.com/cdktn-io/cdktn-aws-go/chimesdkmediapipelines/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference interface {
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
	IssueDetectionConfiguration() AwsMediaInsightsPipelineConfiguration_IssueDetectionConfigurationPropertyOutputReference
	// Experimental.
	IssueDetectionConfigurationInput() *AwsMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty
	// Experimental.
	KeywordMatchConfiguration() AwsMediaInsightsPipelineConfiguration_KeywordMatchConfigurationPropertyOutputReference
	// Experimental.
	KeywordMatchConfigurationInput() *AwsMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty
	// Experimental.
	SentimentConfiguration() AwsMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference
	// Experimental.
	SentimentConfigurationInput() *AwsMediaInsightsPipelineConfiguration_SentimentConfigurationProperty
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
	PutIssueDetectionConfiguration(value *AwsMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty)
	// Experimental.
	PutKeywordMatchConfiguration(value *AwsMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty)
	// Experimental.
	PutSentimentConfiguration(value *AwsMediaInsightsPipelineConfiguration_SentimentConfigurationProperty)
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

// The jsii proxy struct for AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference
type jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) IssueDetectionConfiguration() AwsMediaInsightsPipelineConfiguration_IssueDetectionConfigurationPropertyOutputReference {
	var returns AwsMediaInsightsPipelineConfiguration_IssueDetectionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"issueDetectionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) IssueDetectionConfigurationInput() *AwsMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty {
	var returns *AwsMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty
	_jsii_.Get(
		j,
		"issueDetectionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) KeywordMatchConfiguration() AwsMediaInsightsPipelineConfiguration_KeywordMatchConfigurationPropertyOutputReference {
	var returns AwsMediaInsightsPipelineConfiguration_KeywordMatchConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"keywordMatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) KeywordMatchConfigurationInput() *AwsMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty {
	var returns *AwsMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty
	_jsii_.Get(
		j,
		"keywordMatchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) SentimentConfiguration() AwsMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference {
	var returns AwsMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"sentimentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) SentimentConfigurationInput() *AwsMediaInsightsPipelineConfiguration_SentimentConfigurationProperty {
	var returns *AwsMediaInsightsPipelineConfiguration_SentimentConfigurationProperty
	_jsii_.Get(
		j,
		"sentimentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.RulesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference_Override(a AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.RulesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) PutIssueDetectionConfiguration(value *AwsMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty) {
	if err := a.validatePutIssueDetectionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIssueDetectionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) PutKeywordMatchConfiguration(value *AwsMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty) {
	if err := a.validatePutKeywordMatchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKeywordMatchConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) PutSentimentConfiguration(value *AwsMediaInsightsPipelineConfiguration_SentimentConfigurationProperty) {
	if err := a.validatePutSentimentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSentimentConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ResetIssueDetectionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetIssueDetectionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ResetKeywordMatchConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetKeywordMatchConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ResetSentimentConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSentimentConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

