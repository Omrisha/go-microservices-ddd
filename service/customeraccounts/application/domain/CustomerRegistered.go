package domain

import (
	"go-ddd-services/service/customeraccounts/application/domain/customer/value"
	"go-ddd-services/service/shared/es"
)

type CustomerRegistered struct {
	customerID       value.CustomerID
	emailAddress     value.EmailAddress
	confirmationHash value.ConfirmationHash
	personName       value.PersonName
	meta             es.EventMeta
}

func BuildCustomerRegistered(
	customerId value.CustomerID,
	emailAddress value.EmailAddress,
	confirmationHash value.ConfirmationHash,
	personName value.PersonName,
	streamVersion uint,
) CustomerRegistered {
	event := CustomerRegistered{
		customerID:       customerId,
		emailAddress:     emailAddress,
		confirmationHash: confirmationHash,
		personName:       personName,
	}

	event.meta = es.BuildEventMeta(event, streamVersion)

	return event
}

func (customer CustomerRegistered) CustomerID() value.CustomerID {
	return customer.customerID
}

func (customer CustomerRegistered) EmailAddress() value.EmailAddress {
	return customer.emailAddress
}

func (customer CustomerRegistered) ConfirmationHash() value.ConfirmationHash {
	return customer.confirmationHash
}

func (customer CustomerRegistered) PersonName() value.PersonName {
	return customer.personName
}

func (customer CustomerRegistered) Meta() es.EventMeta {
	return customer.meta
}

func (customer CustomerRegistered) IsFailureEvent() bool {
	return false
}

func (customer CustomerRegistered) FailureReason() error {
	return nil
}
