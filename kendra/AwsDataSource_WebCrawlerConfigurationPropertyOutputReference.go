package kendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kendra/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kendra/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDataSource_WebCrawlerConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthenticationConfiguration() AwsDataSource_AuthenticationConfigurationPropertyOutputReference
	// Experimental.
	AuthenticationConfigurationInput() *AwsDataSource_AuthenticationConfigurationProperty
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
	CrawlDepth() *float64
	// Experimental.
	SetCrawlDepth(val *float64)
	// Experimental.
	CrawlDepthInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDataSource_WebCrawlerConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsDataSource_WebCrawlerConfigurationProperty)
	// Experimental.
	MaxContentSizePerPageInMegaBytes() *float64
	// Experimental.
	SetMaxContentSizePerPageInMegaBytes(val *float64)
	// Experimental.
	MaxContentSizePerPageInMegaBytesInput() *float64
	// Experimental.
	MaxLinksPerPage() *float64
	// Experimental.
	SetMaxLinksPerPage(val *float64)
	// Experimental.
	MaxLinksPerPageInput() *float64
	// Experimental.
	MaxUrlsPerMinuteCrawlRate() *float64
	// Experimental.
	SetMaxUrlsPerMinuteCrawlRate(val *float64)
	// Experimental.
	MaxUrlsPerMinuteCrawlRateInput() *float64
	// Experimental.
	ProxyConfiguration() AwsDataSource_ProxyConfigurationPropertyOutputReference
	// Experimental.
	ProxyConfigurationInput() *AwsDataSource_ProxyConfigurationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UrlExclusionPatterns() *[]*string
	// Experimental.
	SetUrlExclusionPatterns(val *[]*string)
	// Experimental.
	UrlExclusionPatternsInput() *[]*string
	// Experimental.
	UrlInclusionPatterns() *[]*string
	// Experimental.
	SetUrlInclusionPatterns(val *[]*string)
	// Experimental.
	UrlInclusionPatternsInput() *[]*string
	// Experimental.
	Urls() AwsDataSource_UrlsPropertyOutputReference
	// Experimental.
	UrlsInput() *AwsDataSource_UrlsProperty
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
	PutAuthenticationConfiguration(value *AwsDataSource_AuthenticationConfigurationProperty)
	// Experimental.
	PutProxyConfiguration(value *AwsDataSource_ProxyConfigurationProperty)
	// Experimental.
	PutUrls(value *AwsDataSource_UrlsProperty)
	// Experimental.
	ResetAuthenticationConfiguration()
	// Experimental.
	ResetCrawlDepth()
	// Experimental.
	ResetMaxContentSizePerPageInMegaBytes()
	// Experimental.
	ResetMaxLinksPerPage()
	// Experimental.
	ResetMaxUrlsPerMinuteCrawlRate()
	// Experimental.
	ResetProxyConfiguration()
	// Experimental.
	ResetUrlExclusionPatterns()
	// Experimental.
	ResetUrlInclusionPatterns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDataSource_WebCrawlerConfigurationPropertyOutputReference
type jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) AuthenticationConfiguration() AwsDataSource_AuthenticationConfigurationPropertyOutputReference {
	var returns AwsDataSource_AuthenticationConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"authenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) AuthenticationConfigurationInput() *AwsDataSource_AuthenticationConfigurationProperty {
	var returns *AwsDataSource_AuthenticationConfigurationProperty
	_jsii_.Get(
		j,
		"authenticationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) CrawlDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"crawlDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) CrawlDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"crawlDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) InternalValue() *AwsDataSource_WebCrawlerConfigurationProperty {
	var returns *AwsDataSource_WebCrawlerConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) MaxContentSizePerPageInMegaBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxContentSizePerPageInMegaBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) MaxContentSizePerPageInMegaBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxContentSizePerPageInMegaBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) MaxLinksPerPage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLinksPerPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) MaxLinksPerPageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLinksPerPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) MaxUrlsPerMinuteCrawlRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUrlsPerMinuteCrawlRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) MaxUrlsPerMinuteCrawlRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUrlsPerMinuteCrawlRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) ProxyConfiguration() AwsDataSource_ProxyConfigurationPropertyOutputReference {
	var returns AwsDataSource_ProxyConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"proxyConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) ProxyConfigurationInput() *AwsDataSource_ProxyConfigurationProperty {
	var returns *AwsDataSource_ProxyConfigurationProperty
	_jsii_.Get(
		j,
		"proxyConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) UrlExclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlExclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) UrlExclusionPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlExclusionPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) UrlInclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlInclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) UrlInclusionPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlInclusionPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) Urls() AwsDataSource_UrlsPropertyOutputReference {
	var returns AwsDataSource_UrlsPropertyOutputReference
	_jsii_.Get(
		j,
		"urls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) UrlsInput() *AwsDataSource_UrlsProperty {
	var returns *AwsDataSource_UrlsProperty
	_jsii_.Get(
		j,
		"urlsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDataSource_WebCrawlerConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDataSource_WebCrawlerConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDataSource_WebCrawlerConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kendra.AwsDataSource.WebCrawlerConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDataSource_WebCrawlerConfigurationPropertyOutputReference_Override(a AwsDataSource_WebCrawlerConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kendra.AwsDataSource.WebCrawlerConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference)SetCrawlDepth(val *float64) {
	if err := j.validateSetCrawlDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlDepth",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference)SetInternalValue(val *AwsDataSource_WebCrawlerConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference)SetMaxContentSizePerPageInMegaBytes(val *float64) {
	if err := j.validateSetMaxContentSizePerPageInMegaBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxContentSizePerPageInMegaBytes",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference)SetMaxLinksPerPage(val *float64) {
	if err := j.validateSetMaxLinksPerPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLinksPerPage",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference)SetMaxUrlsPerMinuteCrawlRate(val *float64) {
	if err := j.validateSetMaxUrlsPerMinuteCrawlRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUrlsPerMinuteCrawlRate",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference)SetUrlExclusionPatterns(val *[]*string) {
	if err := j.validateSetUrlExclusionPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlExclusionPatterns",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference)SetUrlInclusionPatterns(val *[]*string) {
	if err := j.validateSetUrlInclusionPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlInclusionPatterns",
		val,
	)
}

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) PutAuthenticationConfiguration(value *AwsDataSource_AuthenticationConfigurationProperty) {
	if err := a.validatePutAuthenticationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthenticationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) PutProxyConfiguration(value *AwsDataSource_ProxyConfigurationProperty) {
	if err := a.validatePutProxyConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProxyConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) PutUrls(value *AwsDataSource_UrlsProperty) {
	if err := a.validatePutUrlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUrls",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) ResetAuthenticationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) ResetCrawlDepth() {
	_jsii_.InvokeVoid(
		a,
		"resetCrawlDepth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) ResetMaxContentSizePerPageInMegaBytes() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxContentSizePerPageInMegaBytes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) ResetMaxLinksPerPage() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxLinksPerPage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) ResetMaxUrlsPerMinuteCrawlRate() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxUrlsPerMinuteCrawlRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) ResetProxyConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetProxyConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) ResetUrlExclusionPatterns() {
	_jsii_.InvokeVoid(
		a,
		"resetUrlExclusionPatterns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) ResetUrlInclusionPatterns() {
	_jsii_.InvokeVoid(
		a,
		"resetUrlInclusionPatterns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDataSource_WebCrawlerConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

