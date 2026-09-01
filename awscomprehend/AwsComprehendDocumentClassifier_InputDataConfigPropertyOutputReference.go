package awscomprehend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscomprehend/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscomprehend/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AugmentedManifests() AwsComprehendDocumentClassifier_AugmentedManifestsPropertyList
	// Experimental.
	AugmentedManifestsInput() interface{}
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
	DataFormat() *string
	// Experimental.
	SetDataFormat(val *string)
	// Experimental.
	DataFormatInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsComprehendDocumentClassifier_InputDataConfigProperty
	// Experimental.
	SetInternalValue(val *AwsComprehendDocumentClassifier_InputDataConfigProperty)
	// Experimental.
	LabelDelimiter() *string
	// Experimental.
	SetLabelDelimiter(val *string)
	// Experimental.
	LabelDelimiterInput() *string
	// Experimental.
	S3Uri() *string
	// Experimental.
	SetS3Uri(val *string)
	// Experimental.
	S3UriInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TestS3Uri() *string
	// Experimental.
	SetTestS3Uri(val *string)
	// Experimental.
	TestS3UriInput() *string
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
	PutAugmentedManifests(value interface{})
	// Experimental.
	ResetAugmentedManifests()
	// Experimental.
	ResetDataFormat()
	// Experimental.
	ResetLabelDelimiter()
	// Experimental.
	ResetS3Uri()
	// Experimental.
	ResetTestS3Uri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference
type jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) AugmentedManifests() AwsComprehendDocumentClassifier_AugmentedManifestsPropertyList {
	var returns AwsComprehendDocumentClassifier_AugmentedManifestsPropertyList
	_jsii_.Get(
		j,
		"augmentedManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) AugmentedManifestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"augmentedManifestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) DataFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) InternalValue() *AwsComprehendDocumentClassifier_InputDataConfigProperty {
	var returns *AwsComprehendDocumentClassifier_InputDataConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) LabelDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) LabelDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) S3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) S3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) TestS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) TestS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testS3UriInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-comprehend.AwsComprehendDocumentClassifier.InputDataConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference_Override(a AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-comprehend.AwsComprehendDocumentClassifier.InputDataConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference)SetDataFormat(val *string) {
	if err := j.validateSetDataFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFormat",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference)SetInternalValue(val *AwsComprehendDocumentClassifier_InputDataConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference)SetLabelDelimiter(val *string) {
	if err := j.validateSetLabelDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelDelimiter",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference)SetS3Uri(val *string) {
	if err := j.validateSetS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Uri",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference)SetTestS3Uri(val *string) {
	if err := j.validateSetTestS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"testS3Uri",
		val,
	)
}

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) PutAugmentedManifests(value interface{}) {
	if err := a.validatePutAugmentedManifestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAugmentedManifests",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) ResetAugmentedManifests() {
	_jsii_.InvokeVoid(
		a,
		"resetAugmentedManifests",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) ResetDataFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetDataFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) ResetLabelDelimiter() {
	_jsii_.InvokeVoid(
		a,
		"resetLabelDelimiter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) ResetS3Uri() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Uri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) ResetTestS3Uri() {
	_jsii_.InvokeVoid(
		a,
		"resetTestS3Uri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsComprehendDocumentClassifier_InputDataConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

