//go:build no_runtime_type_checking

package awskms

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfKey_ReplicaKeysPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataTfKey_ReplicaKeysPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfKey_ReplicaKeysPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfKey_ReplicaKeysPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfKey_ReplicaKeysPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfKey_ReplicaKeysPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfKey_ReplicaKeysPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

