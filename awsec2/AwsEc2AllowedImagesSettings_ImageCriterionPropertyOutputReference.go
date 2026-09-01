package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference interface {
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
	CreationDateCondition() AwsEc2AllowedImagesSettings_CreationDateConditionPropertyList
	// Experimental.
	CreationDateConditionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DeprecationTimeCondition() AwsEc2AllowedImagesSettings_DeprecationTimeConditionPropertyList
	// Experimental.
	DeprecationTimeConditionInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	ImageNames() *[]*string
	// Experimental.
	SetImageNames(val *[]*string)
	// Experimental.
	ImageNamesInput() *[]*string
	// Experimental.
	ImageProviders() *[]*string
	// Experimental.
	SetImageProviders(val *[]*string)
	// Experimental.
	ImageProvidersInput() *[]*string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MarketplaceProductCodes() *[]*string
	// Experimental.
	SetMarketplaceProductCodes(val *[]*string)
	// Experimental.
	MarketplaceProductCodesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutCreationDateCondition(value interface{})
	// Experimental.
	PutDeprecationTimeCondition(value interface{})
	// Experimental.
	ResetCreationDateCondition()
	// Experimental.
	ResetDeprecationTimeCondition()
	// Experimental.
	ResetImageNames()
	// Experimental.
	ResetImageProviders()
	// Experimental.
	ResetMarketplaceProductCodes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference
type jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) CreationDateCondition() AwsEc2AllowedImagesSettings_CreationDateConditionPropertyList {
	var returns AwsEc2AllowedImagesSettings_CreationDateConditionPropertyList
	_jsii_.Get(
		j,
		"creationDateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) CreationDateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"creationDateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) DeprecationTimeCondition() AwsEc2AllowedImagesSettings_DeprecationTimeConditionPropertyList {
	var returns AwsEc2AllowedImagesSettings_DeprecationTimeConditionPropertyList
	_jsii_.Get(
		j,
		"deprecationTimeCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) DeprecationTimeConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deprecationTimeConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) ImageNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imageNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) ImageNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imageNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) ImageProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imageProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) ImageProvidersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imageProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) MarketplaceProductCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"marketplaceProductCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) MarketplaceProductCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"marketplaceProductCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsEc2AllowedImagesSettings.ImageCriterionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference_Override(a AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsEc2AllowedImagesSettings.ImageCriterionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference)SetImageNames(val *[]*string) {
	if err := j.validateSetImageNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageNames",
		val,
	)
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference)SetImageProviders(val *[]*string) {
	if err := j.validateSetImageProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageProviders",
		val,
	)
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference)SetMarketplaceProductCodes(val *[]*string) {
	if err := j.validateSetMarketplaceProductCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"marketplaceProductCodes",
		val,
	)
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) PutCreationDateCondition(value interface{}) {
	if err := a.validatePutCreationDateConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCreationDateCondition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) PutDeprecationTimeCondition(value interface{}) {
	if err := a.validatePutDeprecationTimeConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeprecationTimeCondition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) ResetCreationDateCondition() {
	_jsii_.InvokeVoid(
		a,
		"resetCreationDateCondition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) ResetDeprecationTimeCondition() {
	_jsii_.InvokeVoid(
		a,
		"resetDeprecationTimeCondition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) ResetImageNames() {
	_jsii_.InvokeVoid(
		a,
		"resetImageNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) ResetImageProviders() {
	_jsii_.InvokeVoid(
		a,
		"resetImageProviders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) ResetMarketplaceProductCodes() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketplaceProductCodes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEc2AllowedImagesSettings_ImageCriterionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

