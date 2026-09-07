//go:build no_runtime_type_checking

package dynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsGlobalSecondaryIndex_KeySchemaPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsGlobalSecondaryIndex_KeySchemaPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsGlobalSecondaryIndex_KeySchemaPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsGlobalSecondaryIndex_KeySchemaPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsGlobalSecondaryIndex_KeySchemaPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsGlobalSecondaryIndex_KeySchemaPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsGlobalSecondaryIndex_KeySchemaPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsGlobalSecondaryIndex_KeySchemaPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

