package awschimesdkmediapipelines

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awschimesdkmediapipelines/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awschimesdkmediapipelines/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference interface {
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
	IssueDetectionConfiguration() TfMediaInsightsPipelineConfiguration_IssueDetectionConfigurationPropertyOutputReference
	// Experimental.
	IssueDetectionConfigurationInput() *TfMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty
	// Experimental.
	KeywordMatchConfiguration() TfMediaInsightsPipelineConfiguration_KeywordMatchConfigurationPropertyOutputReference
	// Experimental.
	KeywordMatchConfigurationInput() *TfMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty
	// Experimental.
	SentimentConfiguration() TfMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference
	// Experimental.
	SentimentConfigurationInput() *TfMediaInsightsPipelineConfiguration_SentimentConfigurationProperty
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
	PutIssueDetectionConfiguration(value *TfMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty)
	// Experimental.
	PutKeywordMatchConfiguration(value *TfMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty)
	// Experimental.
	PutSentimentConfiguration(value *TfMediaInsightsPipelineConfiguration_SentimentConfigurationProperty)
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

// The jsii proxy struct for TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference
type jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) IssueDetectionConfiguration() TfMediaInsightsPipelineConfiguration_IssueDetectionConfigurationPropertyOutputReference {
	var returns TfMediaInsightsPipelineConfiguration_IssueDetectionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"issueDetectionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) IssueDetectionConfigurationInput() *TfMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty {
	var returns *TfMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty
	_jsii_.Get(
		j,
		"issueDetectionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) KeywordMatchConfiguration() TfMediaInsightsPipelineConfiguration_KeywordMatchConfigurationPropertyOutputReference {
	var returns TfMediaInsightsPipelineConfiguration_KeywordMatchConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"keywordMatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) KeywordMatchConfigurationInput() *TfMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty {
	var returns *TfMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty
	_jsii_.Get(
		j,
		"keywordMatchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) SentimentConfiguration() TfMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference {
	var returns TfMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"sentimentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) SentimentConfigurationInput() *TfMediaInsightsPipelineConfiguration_SentimentConfigurationProperty {
	var returns *TfMediaInsightsPipelineConfiguration_SentimentConfigurationProperty
	_jsii_.Get(
		j,
		"sentimentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMediaInsightsPipelineConfiguration_RulesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.TfMediaInsightsPipelineConfiguration.RulesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference_Override(t TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.TfMediaInsightsPipelineConfiguration.RulesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) PutIssueDetectionConfiguration(value *TfMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty) {
	if err := t.validatePutIssueDetectionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIssueDetectionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) PutKeywordMatchConfiguration(value *TfMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty) {
	if err := t.validatePutKeywordMatchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKeywordMatchConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) PutSentimentConfiguration(value *TfMediaInsightsPipelineConfiguration_SentimentConfigurationProperty) {
	if err := t.validatePutSentimentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSentimentConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ResetIssueDetectionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetIssueDetectionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ResetKeywordMatchConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetKeywordMatchConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ResetSentimentConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSentimentConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_RulesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

