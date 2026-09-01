package awschimesdkmediapipelines

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awschimesdkmediapipelines/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awschimesdkmediapipelines/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference interface {
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
	InternalValue() *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationProperty)
	// Experimental.
	RuleName() *string
	// Experimental.
	SetRuleName(val *string)
	// Experimental.
	RuleNameInput() *string
	// Experimental.
	SentimentType() *string
	// Experimental.
	SetSentimentType(val *string)
	// Experimental.
	SentimentTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimePeriod() *float64
	// Experimental.
	SetTimePeriod(val *float64)
	// Experimental.
	TimePeriodInput() *float64
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

// The jsii proxy struct for AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference
type jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) InternalValue() *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationProperty {
	var returns *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) RuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) SentimentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sentimentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) SentimentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sentimentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) TimePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) TimePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timePeriodInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration.SentimentConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference_Override(a AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration.SentimentConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference)SetInternalValue(val *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference)SetRuleName(val *string) {
	if err := j.validateSetRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference)SetSentimentType(val *string) {
	if err := j.validateSetSentimentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sentimentType",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference)SetTimePeriod(val *float64) {
	if err := j.validateSetTimePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timePeriod",
		val,
	)
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

