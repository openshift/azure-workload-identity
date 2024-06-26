package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryIndexOperation 
type EdiscoveryIndexOperation struct {
    CaseOperation
}
// NewEdiscoveryIndexOperation instantiates a new EdiscoveryIndexOperation and sets the default values.
func NewEdiscoveryIndexOperation()(*EdiscoveryIndexOperation) {
    m := &EdiscoveryIndexOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateEdiscoveryIndexOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryIndexOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryIndexOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryIndexOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *EdiscoveryIndexOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// EdiscoveryIndexOperationable 
type EdiscoveryIndexOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
