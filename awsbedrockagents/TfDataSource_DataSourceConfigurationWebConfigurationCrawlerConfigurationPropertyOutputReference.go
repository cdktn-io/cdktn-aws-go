package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference interface {
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
	CrawlerLimits() TfDataSource_CrawlerLimitsPropertyList
	// Experimental.
	CrawlerLimitsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	ExclusionFilters() *[]*string
	// Experimental.
	SetExclusionFilters(val *[]*string)
	// Experimental.
	ExclusionFiltersInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InclusionFilters() *[]*string
	// Experimental.
	SetInclusionFilters(val *[]*string)
	// Experimental.
	InclusionFiltersInput() *[]*string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Scope() *string
	// Experimental.
	SetScope(val *string)
	// Experimental.
	ScopeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UserAgent() *string
	// Experimental.
	SetUserAgent(val *string)
	// Experimental.
	UserAgentInput() *string
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
	PutCrawlerLimits(value interface{})
	// Experimental.
	ResetCrawlerLimits()
	// Experimental.
	ResetExclusionFilters()
	// Experimental.
	ResetInclusionFilters()
	// Experimental.
	ResetScope()
	// Experimental.
	ResetUserAgent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference
type jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) CrawlerLimits() TfDataSource_CrawlerLimitsPropertyList {
	var returns TfDataSource_CrawlerLimitsPropertyList
	_jsii_.Get(
		j,
		"crawlerLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) CrawlerLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ExclusionFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ExclusionFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) InclusionFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) InclusionFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) UserAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) UserAgentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgentInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfDataSource.DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference_Override(t TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfDataSource.DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetExclusionFilters(val *[]*string) {
	if err := j.validateSetExclusionFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusionFilters",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetInclusionFilters(val *[]*string) {
	if err := j.validateSetInclusionFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inclusionFilters",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetUserAgent(val *string) {
	if err := j.validateSetUserAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userAgent",
		val,
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) PutCrawlerLimits(value interface{}) {
	if err := t.validatePutCrawlerLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCrawlerLimits",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ResetCrawlerLimits() {
	_jsii_.InvokeVoid(
		t,
		"resetCrawlerLimits",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ResetExclusionFilters() {
	_jsii_.InvokeVoid(
		t,
		"resetExclusionFilters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ResetInclusionFilters() {
	_jsii_.InvokeVoid(
		t,
		"resetInclusionFilters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		t,
		"resetScope",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ResetUserAgent() {
	_jsii_.InvokeVoid(
		t,
		"resetUserAgent",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

