package awskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskendra/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskendra/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataSource_ConfigurationPropertyOutputReference interface {
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
	InternalValue() *TfDataSource_ConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfDataSource_ConfigurationProperty)
	// Experimental.
	S3Configuration() TfDataSource_S3ConfigurationPropertyOutputReference
	// Experimental.
	S3ConfigurationInput() *TfDataSource_S3ConfigurationProperty
	// Experimental.
	TemplateConfiguration() TfDataSource_TemplateConfigurationPropertyOutputReference
	// Experimental.
	TemplateConfigurationInput() *TfDataSource_TemplateConfigurationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WebCrawlerConfiguration() TfDataSource_WebCrawlerConfigurationPropertyOutputReference
	// Experimental.
	WebCrawlerConfigurationInput() *TfDataSource_WebCrawlerConfigurationProperty
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
	PutS3Configuration(value *TfDataSource_S3ConfigurationProperty)
	// Experimental.
	PutTemplateConfiguration(value *TfDataSource_TemplateConfigurationProperty)
	// Experimental.
	PutWebCrawlerConfiguration(value *TfDataSource_WebCrawlerConfigurationProperty)
	// Experimental.
	ResetS3Configuration()
	// Experimental.
	ResetTemplateConfiguration()
	// Experimental.
	ResetWebCrawlerConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataSource_ConfigurationPropertyOutputReference
type jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) InternalValue() *TfDataSource_ConfigurationProperty {
	var returns *TfDataSource_ConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) S3Configuration() TfDataSource_S3ConfigurationPropertyOutputReference {
	var returns TfDataSource_S3ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) S3ConfigurationInput() *TfDataSource_S3ConfigurationProperty {
	var returns *TfDataSource_S3ConfigurationProperty
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) TemplateConfiguration() TfDataSource_TemplateConfigurationPropertyOutputReference {
	var returns TfDataSource_TemplateConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"templateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) TemplateConfigurationInput() *TfDataSource_TemplateConfigurationProperty {
	var returns *TfDataSource_TemplateConfigurationProperty
	_jsii_.Get(
		j,
		"templateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) WebCrawlerConfiguration() TfDataSource_WebCrawlerConfigurationPropertyOutputReference {
	var returns TfDataSource_WebCrawlerConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"webCrawlerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) WebCrawlerConfigurationInput() *TfDataSource_WebCrawlerConfigurationProperty {
	var returns *TfDataSource_WebCrawlerConfigurationProperty
	_jsii_.Get(
		j,
		"webCrawlerConfigurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataSource_ConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDataSource_ConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataSource_ConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kendra.TfDataSource.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataSource_ConfigurationPropertyOutputReference_Override(t TfDataSource_ConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kendra.TfDataSource.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference)SetInternalValue(val *TfDataSource_ConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) PutS3Configuration(value *TfDataSource_S3ConfigurationProperty) {
	if err := t.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) PutTemplateConfiguration(value *TfDataSource_TemplateConfigurationProperty) {
	if err := t.validatePutTemplateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTemplateConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) PutWebCrawlerConfiguration(value *TfDataSource_WebCrawlerConfigurationProperty) {
	if err := t.validatePutWebCrawlerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWebCrawlerConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) ResetS3Configuration() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Configuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) ResetTemplateConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetTemplateConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) ResetWebCrawlerConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetWebCrawlerConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataSource_ConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

