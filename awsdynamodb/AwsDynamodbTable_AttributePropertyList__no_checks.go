//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsDynamodbTable_AttributePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsDynamodbTable_AttributePropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsDynamodbTable_AttributePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsDynamodbTable_AttributePropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsDynamodbTable_AttributePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsDynamodbTable_AttributePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsDynamodbTable_AttributePropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsDynamodbTable_AttributePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

