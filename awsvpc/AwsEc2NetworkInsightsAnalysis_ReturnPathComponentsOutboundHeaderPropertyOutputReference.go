package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference interface {
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
	DestinationAddresses() *[]*string
	// Experimental.
	DestinationPortRanges() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderDestinationPortRangesPropertyList
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderProperty
	// Experimental.
	SetInternalValue(val *AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderProperty)
	// Experimental.
	Protocol() *string
	// Experimental.
	SourceAddresses() *[]*string
	// Experimental.
	SourcePortRanges() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderSourcePortRangesPropertyList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference
type jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) DestinationAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) DestinationPortRanges() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderDestinationPortRangesPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderDestinationPortRangesPropertyList
	_jsii_.Get(
		j,
		"destinationPortRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) InternalValue() *AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderProperty {
	var returns *AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) SourceAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) SourcePortRanges() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderSourcePortRangesPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderSourcePortRangesPropertyList
	_jsii_.Get(
		j,
		"sourcePortRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsEc2NetworkInsightsAnalysis.ReturnPathComponentsOutboundHeaderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference_Override(a AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsEc2NetworkInsightsAnalysis.ReturnPathComponentsOutboundHeaderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference)SetInternalValue(val *AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

