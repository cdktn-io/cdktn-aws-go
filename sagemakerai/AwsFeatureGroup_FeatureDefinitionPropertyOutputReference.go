package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFeatureGroup_FeatureDefinitionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CollectionConfig() AwsFeatureGroup_CollectionConfigPropertyOutputReference
	// Experimental.
	CollectionConfigInput() *AwsFeatureGroup_CollectionConfigProperty
	// Experimental.
	CollectionType() *string
	// Experimental.
	SetCollectionType(val *string)
	// Experimental.
	CollectionTypeInput() *string
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
	FeatureName() *string
	// Experimental.
	SetFeatureName(val *string)
	// Experimental.
	FeatureNameInput() *string
	// Experimental.
	FeatureType() *string
	// Experimental.
	SetFeatureType(val *string)
	// Experimental.
	FeatureTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
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
	PutCollectionConfig(value *AwsFeatureGroup_CollectionConfigProperty)
	// Experimental.
	ResetCollectionConfig()
	// Experimental.
	ResetCollectionType()
	// Experimental.
	ResetFeatureName()
	// Experimental.
	ResetFeatureType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsFeatureGroup_FeatureDefinitionPropertyOutputReference
type jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) CollectionConfig() AwsFeatureGroup_CollectionConfigPropertyOutputReference {
	var returns AwsFeatureGroup_CollectionConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"collectionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) CollectionConfigInput() *AwsFeatureGroup_CollectionConfigProperty {
	var returns *AwsFeatureGroup_CollectionConfigProperty
	_jsii_.Get(
		j,
		"collectionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) CollectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) CollectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) FeatureName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) FeatureNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) FeatureType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) FeatureTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFeatureGroup_FeatureDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsFeatureGroup_FeatureDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFeatureGroup_FeatureDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsFeatureGroup.FeatureDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFeatureGroup_FeatureDefinitionPropertyOutputReference_Override(a AwsFeatureGroup_FeatureDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsFeatureGroup.FeatureDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference)SetCollectionType(val *string) {
	if err := j.validateSetCollectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionType",
		val,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference)SetFeatureName(val *string) {
	if err := j.validateSetFeatureNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featureName",
		val,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference)SetFeatureType(val *string) {
	if err := j.validateSetFeatureTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featureType",
		val,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) PutCollectionConfig(value *AwsFeatureGroup_CollectionConfigProperty) {
	if err := a.validatePutCollectionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCollectionConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) ResetCollectionConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCollectionConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) ResetCollectionType() {
	_jsii_.InvokeVoid(
		a,
		"resetCollectionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) ResetFeatureName() {
	_jsii_.InvokeVoid(
		a,
		"resetFeatureName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) ResetFeatureType() {
	_jsii_.InvokeVoid(
		a,
		"resetFeatureType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFeatureGroup_FeatureDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

