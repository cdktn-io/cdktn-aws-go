package bedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference interface {
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
	CrawlerLimits() AwsDataSource_CrawlerLimitsPropertyList
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

// The jsii proxy struct for AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference
type jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) CrawlerLimits() AwsDataSource_CrawlerLimitsPropertyList {
	var returns AwsDataSource_CrawlerLimitsPropertyList
	_jsii_.Get(
		j,
		"crawlerLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) CrawlerLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ExclusionFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ExclusionFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) InclusionFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) InclusionFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) UserAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) UserAgentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgentInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsDataSource.DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference_Override(a AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsDataSource.DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetExclusionFilters(val *[]*string) {
	if err := j.validateSetExclusionFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusionFilters",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetInclusionFilters(val *[]*string) {
	if err := j.validateSetInclusionFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inclusionFilters",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference)SetUserAgent(val *string) {
	if err := j.validateSetUserAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userAgent",
		val,
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) PutCrawlerLimits(value interface{}) {
	if err := a.validatePutCrawlerLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCrawlerLimits",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ResetCrawlerLimits() {
	_jsii_.InvokeVoid(
		a,
		"resetCrawlerLimits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ResetExclusionFilters() {
	_jsii_.InvokeVoid(
		a,
		"resetExclusionFilters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ResetInclusionFilters() {
	_jsii_.InvokeVoid(
		a,
		"resetInclusionFilters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		a,
		"resetScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ResetUserAgent() {
	_jsii_.InvokeVoid(
		a,
		"resetUserAgent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

