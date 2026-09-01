package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference interface {
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
	SetDestinationAddresses(val *[]*string)
	// Experimental.
	DestinationAddressesInput() *[]*string
	// Experimental.
	DestinationPorts() *[]*string
	// Experimental.
	SetDestinationPorts(val *[]*string)
	// Experimental.
	DestinationPortsInput() *[]*string
	// Experimental.
	DestinationPrefixLists() *[]*string
	// Experimental.
	SetDestinationPrefixLists(val *[]*string)
	// Experimental.
	DestinationPrefixListsInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Protocols() *[]*string
	// Experimental.
	SetProtocols(val *[]*string)
	// Experimental.
	ProtocolsInput() *[]*string
	// Experimental.
	SourceAddresses() *[]*string
	// Experimental.
	SetSourceAddresses(val *[]*string)
	// Experimental.
	SourceAddressesInput() *[]*string
	// Experimental.
	SourcePorts() *[]*string
	// Experimental.
	SetSourcePorts(val *[]*string)
	// Experimental.
	SourcePortsInput() *[]*string
	// Experimental.
	SourcePrefixLists() *[]*string
	// Experimental.
	SetSourcePrefixLists(val *[]*string)
	// Experimental.
	SourcePrefixListsInput() *[]*string
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
	ResetDestinationAddresses()
	// Experimental.
	ResetDestinationPorts()
	// Experimental.
	ResetDestinationPrefixLists()
	// Experimental.
	ResetProtocols()
	// Experimental.
	ResetSourceAddresses()
	// Experimental.
	ResetSourcePorts()
	// Experimental.
	ResetSourcePrefixLists()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference
type jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) DestinationAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) DestinationAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) DestinationPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) DestinationPortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) DestinationPrefixLists() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPrefixLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) DestinationPrefixListsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPrefixListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) SourceAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) SourceAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) SourcePorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) SourcePortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) SourcePrefixLists() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePrefixLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) SourcePrefixListsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePrefixListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsEc2NetworkInsightsAccessScope.MatchPathsSourcePacketHeaderStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference_Override(a AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsEc2NetworkInsightsAccessScope.MatchPathsSourcePacketHeaderStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference)SetDestinationAddresses(val *[]*string) {
	if err := j.validateSetDestinationAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddresses",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference)SetDestinationPorts(val *[]*string) {
	if err := j.validateSetDestinationPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPorts",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference)SetDestinationPrefixLists(val *[]*string) {
	if err := j.validateSetDestinationPrefixListsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPrefixLists",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference)SetSourceAddresses(val *[]*string) {
	if err := j.validateSetSourceAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddresses",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference)SetSourcePorts(val *[]*string) {
	if err := j.validateSetSourcePortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePorts",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference)SetSourcePrefixLists(val *[]*string) {
	if err := j.validateSetSourcePrefixListsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePrefixLists",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) ResetDestinationAddresses() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationAddresses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) ResetDestinationPorts() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationPorts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) ResetDestinationPrefixLists() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationPrefixLists",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) ResetProtocols() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocols",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) ResetSourceAddresses() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceAddresses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) ResetSourcePorts() {
	_jsii_.InvokeVoid(
		a,
		"resetSourcePorts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) ResetSourcePrefixLists() {
	_jsii_.InvokeVoid(
		a,
		"resetSourcePrefixLists",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAccessScope_MatchPathsSourcePacketHeaderStatementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

