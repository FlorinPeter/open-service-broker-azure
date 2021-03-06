package postgresql

import (
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (d *databaseManager) Bind(
	instance service.Instance,
	_ service.BindingParameters,
) (service.BindingDetails, service.SecureBindingDetails, error) {
	pdt := dbmsInstanceDetails{}
	if err :=
		service.GetStructFromMap(instance.Parent.Details, &pdt); err != nil {
		return nil, nil, err
	}
	spdt := secureDBMSInstanceDetails{}
	if err :=
		service.GetStructFromMap(instance.Parent.SecureDetails, &spdt); err != nil {
		return nil, nil, err
	}
	dt := databaseInstanceDetails{}
	if err := service.GetStructFromMap(instance.Details, &dt); err != nil {
		return nil, nil, err
	}
	bd, spd, err := createBinding(
		isSSLRequired(*instance.Parent.ProvisioningParameters),
		pdt.ServerName,
		spdt.AdministratorLoginPassword,
		pdt.FullyQualifiedDomainName,
		dt.DatabaseName,
	)
	return bd, spd, err
}

func (d *databaseManager) GetCredentials(
	instance service.Instance,
	binding service.Binding,
) (service.Credentials, error) {
	pdt := dbmsInstanceDetails{}
	if err :=
		service.GetStructFromMap(instance.Parent.Details, &pdt); err != nil {
		return nil, err
	}
	dt := databaseInstanceDetails{}
	if err := service.GetStructFromMap(instance.Details, &dt); err != nil {
		return nil, err
	}
	bd := bindingDetails{}
	if err := service.GetStructFromMap(binding.Details, &bd); err != nil {
		return nil, err
	}
	sbd := secureBindingDetails{}
	if err := service.GetStructFromMap(binding.SecureDetails, &sbd); err != nil {
		return nil, err
	}
	cred := createCredential(
		pdt.FullyQualifiedDomainName,
		isSSLRequired(*instance.Parent.ProvisioningParameters),
		pdt.ServerName,
		dt.DatabaseName,
		bd,
		sbd,
	)
	return cred, nil
}
