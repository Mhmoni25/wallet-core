// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Bitcoin.proto

package bitcoin

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Transaction struct {
	// Transaction data format version.
	Version int32 `protobuf:"zigzag32,1,opt,name=version,proto3" json:"version,omitempty"`
	// The block number or timestamp at which this transaction is unlocked.
	LockTime uint32 `protobuf:"varint,2,opt,name=lockTime,proto3" json:"lockTime,omitempty"`
	// A list of 1 or more transaction inputs or sources for coins.
	Inputs []*TransactionInput `protobuf:"bytes,3,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// A list of 1 or more transaction outputs or destinations for coins
	Outputs              []*TransactionOutput `protobuf:"bytes,4,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fbb050c4cb9ab40, []int{0}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Transaction) GetLockTime() uint32 {
	if m != nil {
		return m.LockTime
	}
	return 0
}

func (m *Transaction) GetInputs() []*TransactionInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Transaction) GetOutputs() []*TransactionOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// Bitcoin transaction input.
type TransactionInput struct {
	// Reference to the previous transaction's output.
	PreviousOutput *OutPoint `protobuf:"bytes,1,opt,name=previousOutput,proto3" json:"previousOutput,omitempty"`
	// Transaction version as defined by the sender.
	Sequence uint32 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Computational script for confirming transaction authorization.
	Script               []byte   `protobuf:"bytes,3,opt,name=script,proto3" json:"script,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionInput) Reset()         { *m = TransactionInput{} }
func (m *TransactionInput) String() string { return proto.CompactTextString(m) }
func (*TransactionInput) ProtoMessage()    {}
func (*TransactionInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fbb050c4cb9ab40, []int{1}
}

func (m *TransactionInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionInput.Unmarshal(m, b)
}
func (m *TransactionInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionInput.Marshal(b, m, deterministic)
}
func (m *TransactionInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionInput.Merge(m, src)
}
func (m *TransactionInput) XXX_Size() int {
	return xxx_messageInfo_TransactionInput.Size(m)
}
func (m *TransactionInput) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionInput.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionInput proto.InternalMessageInfo

func (m *TransactionInput) GetPreviousOutput() *OutPoint {
	if m != nil {
		return m.PreviousOutput
	}
	return nil
}

func (m *TransactionInput) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *TransactionInput) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

// Bitcoin transaction out-point reference.
type OutPoint struct {
	// The hash of the referenced transaction.
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// The index of the specific output in the transaction.
	Index uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	// Transaction version as defined by the sender.
	Sequence             uint32   `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutPoint) Reset()         { *m = OutPoint{} }
func (m *OutPoint) String() string { return proto.CompactTextString(m) }
func (*OutPoint) ProtoMessage()    {}
func (*OutPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fbb050c4cb9ab40, []int{2}
}

func (m *OutPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutPoint.Unmarshal(m, b)
}
func (m *OutPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutPoint.Marshal(b, m, deterministic)
}
func (m *OutPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutPoint.Merge(m, src)
}
func (m *OutPoint) XXX_Size() int {
	return xxx_messageInfo_OutPoint.Size(m)
}
func (m *OutPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_OutPoint.DiscardUnknown(m)
}

var xxx_messageInfo_OutPoint proto.InternalMessageInfo

func (m *OutPoint) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *OutPoint) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *OutPoint) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

// Bitcoin transaction output.
type TransactionOutput struct {
	// Transaction amount.
	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	// Usually contains the public key as a Bitcoin script setting up conditions to claim this output.
	Script               []byte   `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionOutput) Reset()         { *m = TransactionOutput{} }
func (m *TransactionOutput) String() string { return proto.CompactTextString(m) }
func (*TransactionOutput) ProtoMessage()    {}
func (*TransactionOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fbb050c4cb9ab40, []int{3}
}

func (m *TransactionOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionOutput.Unmarshal(m, b)
}
func (m *TransactionOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionOutput.Marshal(b, m, deterministic)
}
func (m *TransactionOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionOutput.Merge(m, src)
}
func (m *TransactionOutput) XXX_Size() int {
	return xxx_messageInfo_TransactionOutput.Size(m)
}
func (m *TransactionOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionOutput.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionOutput proto.InternalMessageInfo

func (m *TransactionOutput) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TransactionOutput) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

type UnspentTransaction struct {
	OutPoint             *OutPoint `protobuf:"bytes,1,opt,name=out_point,json=outPoint,proto3" json:"out_point,omitempty"`
	Script               []byte    `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	Amount               int64     `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UnspentTransaction) Reset()         { *m = UnspentTransaction{} }
func (m *UnspentTransaction) String() string { return proto.CompactTextString(m) }
func (*UnspentTransaction) ProtoMessage()    {}
func (*UnspentTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fbb050c4cb9ab40, []int{4}
}

func (m *UnspentTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnspentTransaction.Unmarshal(m, b)
}
func (m *UnspentTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnspentTransaction.Marshal(b, m, deterministic)
}
func (m *UnspentTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnspentTransaction.Merge(m, src)
}
func (m *UnspentTransaction) XXX_Size() int {
	return xxx_messageInfo_UnspentTransaction.Size(m)
}
func (m *UnspentTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_UnspentTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_UnspentTransaction proto.InternalMessageInfo

func (m *UnspentTransaction) GetOutPoint() *OutPoint {
	if m != nil {
		return m.OutPoint
	}
	return nil
}

func (m *UnspentTransaction) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *UnspentTransaction) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// Input data necessary to create a signed transaction.
type SigningInput struct {
	// Hash type to use when signing.
	HashType uint32 `protobuf:"varint,1,opt,name=hash_type,json=hashType,proto3" json:"hash_type,omitempty"`
	// Amount to send.
	Amount int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// Transaction fee per byte.
	ByteFee int64 `protobuf:"varint,3,opt,name=byte_fee,json=byteFee,proto3" json:"byte_fee,omitempty"`
	// Recipient's address.
	ToAddress string `protobuf:"bytes,4,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	// Change address.
	ChangeAddress string `protobuf:"bytes,5,opt,name=change_address,json=changeAddress,proto3" json:"change_address,omitempty"`
	// Available private keys.
	PrivateKey [][]byte `protobuf:"bytes,10,rep,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// Available redeem scripts indexed by script hash.
	Scripts map[string][]byte `protobuf:"bytes,11,rep,name=scripts,proto3" json:"scripts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Available unspent transaction outputs.
	Utxo []*UnspentTransaction `protobuf:"bytes,12,rep,name=utxo,proto3" json:"utxo,omitempty"`
	// If sending max amount.
	UseMaxAmount bool `protobuf:"varint,13,opt,name=use_max_amount,json=useMaxAmount,proto3" json:"use_max_amount,omitempty"`
	// Coin type (forks).
	CoinType uint32 `protobuf:"varint,14,opt,name=coin_type,json=coinType,proto3" json:"coin_type,omitempty"`
	// Optional transaction plan
	Plan                 *TransactionPlan `protobuf:"bytes,15,opt,name=plan,proto3" json:"plan,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SigningInput) Reset()         { *m = SigningInput{} }
func (m *SigningInput) String() string { return proto.CompactTextString(m) }
func (*SigningInput) ProtoMessage()    {}
func (*SigningInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fbb050c4cb9ab40, []int{5}
}

func (m *SigningInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigningInput.Unmarshal(m, b)
}
func (m *SigningInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigningInput.Marshal(b, m, deterministic)
}
func (m *SigningInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigningInput.Merge(m, src)
}
func (m *SigningInput) XXX_Size() int {
	return xxx_messageInfo_SigningInput.Size(m)
}
func (m *SigningInput) XXX_DiscardUnknown() {
	xxx_messageInfo_SigningInput.DiscardUnknown(m)
}

var xxx_messageInfo_SigningInput proto.InternalMessageInfo

func (m *SigningInput) GetHashType() uint32 {
	if m != nil {
		return m.HashType
	}
	return 0
}

func (m *SigningInput) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *SigningInput) GetByteFee() int64 {
	if m != nil {
		return m.ByteFee
	}
	return 0
}

func (m *SigningInput) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *SigningInput) GetChangeAddress() string {
	if m != nil {
		return m.ChangeAddress
	}
	return ""
}

func (m *SigningInput) GetPrivateKey() [][]byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *SigningInput) GetScripts() map[string][]byte {
	if m != nil {
		return m.Scripts
	}
	return nil
}

func (m *SigningInput) GetUtxo() []*UnspentTransaction {
	if m != nil {
		return m.Utxo
	}
	return nil
}

func (m *SigningInput) GetUseMaxAmount() bool {
	if m != nil {
		return m.UseMaxAmount
	}
	return false
}

func (m *SigningInput) GetCoinType() uint32 {
	if m != nil {
		return m.CoinType
	}
	return 0
}

func (m *SigningInput) GetPlan() *TransactionPlan {
	if m != nil {
		return m.Plan
	}
	return nil
}

// Describes a preliminary transaction plan.
type TransactionPlan struct {
	// Amount to be received at the other end.
	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// Maximum available amount.
	AvailableAmount int64 `protobuf:"varint,2,opt,name=available_amount,json=availableAmount,proto3" json:"available_amount,omitempty"`
	// Estimated transaction fee.
	Fee int64 `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
	// Change.
	Change int64 `protobuf:"varint,4,opt,name=change,proto3" json:"change,omitempty"`
	// Selected unspent transaction outputs.
	Utxos []*UnspentTransaction `protobuf:"bytes,5,rep,name=utxos,proto3" json:"utxos,omitempty"`
	// Zcash branch id
	BranchId             []byte   `protobuf:"bytes,6,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionPlan) Reset()         { *m = TransactionPlan{} }
func (m *TransactionPlan) String() string { return proto.CompactTextString(m) }
func (*TransactionPlan) ProtoMessage()    {}
func (*TransactionPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fbb050c4cb9ab40, []int{6}
}

func (m *TransactionPlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionPlan.Unmarshal(m, b)
}
func (m *TransactionPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionPlan.Marshal(b, m, deterministic)
}
func (m *TransactionPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionPlan.Merge(m, src)
}
func (m *TransactionPlan) XXX_Size() int {
	return xxx_messageInfo_TransactionPlan.Size(m)
}
func (m *TransactionPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionPlan.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionPlan proto.InternalMessageInfo

func (m *TransactionPlan) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransactionPlan) GetAvailableAmount() int64 {
	if m != nil {
		return m.AvailableAmount
	}
	return 0
}

func (m *TransactionPlan) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *TransactionPlan) GetChange() int64 {
	if m != nil {
		return m.Change
	}
	return 0
}

func (m *TransactionPlan) GetUtxos() []*UnspentTransaction {
	if m != nil {
		return m.Utxos
	}
	return nil
}

func (m *TransactionPlan) GetBranchId() []byte {
	if m != nil {
		return m.BranchId
	}
	return nil
}

// Transaction signing output.
type SigningOutput struct {
	// Resulting transaction. Note that the amount may be different than the requested amount to account for fees and available funds.
	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	// Signed and encoded transaction bytes.
	Encoded []byte `protobuf:"bytes,2,opt,name=encoded,proto3" json:"encoded,omitempty"`
	// Transaction id
	TransactionId string `protobuf:"bytes,3,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// Optional error message
	Error                string   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SigningOutput) Reset()         { *m = SigningOutput{} }
func (m *SigningOutput) String() string { return proto.CompactTextString(m) }
func (*SigningOutput) ProtoMessage()    {}
func (*SigningOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fbb050c4cb9ab40, []int{7}
}

func (m *SigningOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigningOutput.Unmarshal(m, b)
}
func (m *SigningOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigningOutput.Marshal(b, m, deterministic)
}
func (m *SigningOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigningOutput.Merge(m, src)
}
func (m *SigningOutput) XXX_Size() int {
	return xxx_messageInfo_SigningOutput.Size(m)
}
func (m *SigningOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_SigningOutput.DiscardUnknown(m)
}

var xxx_messageInfo_SigningOutput proto.InternalMessageInfo

func (m *SigningOutput) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *SigningOutput) GetEncoded() []byte {
	if m != nil {
		return m.Encoded
	}
	return nil
}

func (m *SigningOutput) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *SigningOutput) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Transaction)(nil), "TW.Bitcoin.Proto.Transaction")
	proto.RegisterType((*TransactionInput)(nil), "TW.Bitcoin.Proto.TransactionInput")
	proto.RegisterType((*OutPoint)(nil), "TW.Bitcoin.Proto.OutPoint")
	proto.RegisterType((*TransactionOutput)(nil), "TW.Bitcoin.Proto.TransactionOutput")
	proto.RegisterType((*UnspentTransaction)(nil), "TW.Bitcoin.Proto.UnspentTransaction")
	proto.RegisterType((*SigningInput)(nil), "TW.Bitcoin.Proto.SigningInput")
	proto.RegisterMapType((map[string][]byte)(nil), "TW.Bitcoin.Proto.SigningInput.ScriptsEntry")
	proto.RegisterType((*TransactionPlan)(nil), "TW.Bitcoin.Proto.TransactionPlan")
	proto.RegisterType((*SigningOutput)(nil), "TW.Bitcoin.Proto.SigningOutput")
}

func init() { proto.RegisterFile("Bitcoin.proto", fileDescriptor_9fbb050c4cb9ab40) }

var fileDescriptor_9fbb050c4cb9ab40 = []byte{
	// 715 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdb, 0x6a, 0xdb, 0x4c,
	0x10, 0x46, 0x96, 0xe3, 0xc3, 0xf8, 0x10, 0x67, 0xf9, 0x0f, 0xfa, 0x13, 0xc2, 0xef, 0xaa, 0x29,
	0xb8, 0x14, 0x7c, 0x91, 0x52, 0x1a, 0x0c, 0xa5, 0x24, 0x90, 0x42, 0x28, 0x25, 0x66, 0xe3, 0xd2,
	0x4b, 0xb1, 0x96, 0xb6, 0xf1, 0x36, 0xca, 0xae, 0x2a, 0xad, 0x5c, 0xfb, 0xa2, 0x2f, 0xd0, 0x17,
	0xe9, 0x5b, 0xf4, 0x35, 0xfa, 0x32, 0xbd, 0x28, 0x7b, 0x90, 0x23, 0xdb, 0x34, 0xe4, 0xca, 0xfb,
	0x0d, 0x33, 0xdf, 0xec, 0xf7, 0xcd, 0xac, 0x05, 0x9d, 0x33, 0x26, 0x43, 0xc1, 0xf8, 0x30, 0x49,
	0x85, 0x14, 0xa8, 0x37, 0xf9, 0x30, 0x2c, 0x22, 0x63, 0x15, 0xf1, 0x7f, 0x38, 0xd0, 0x9a, 0xa4,
	0x84, 0x67, 0x24, 0x94, 0x4c, 0x70, 0xe4, 0x41, 0x7d, 0x4e, 0xd3, 0x8c, 0x09, 0xee, 0x39, 0x7d,
	0x67, 0xb0, 0x87, 0x0b, 0x88, 0xf6, 0xa1, 0x11, 0x8b, 0xf0, 0x66, 0xc2, 0x6e, 0xa9, 0x57, 0xe9,
	0x3b, 0x83, 0x0e, 0x5e, 0x61, 0x34, 0x82, 0x1a, 0xe3, 0x49, 0x2e, 0x33, 0xcf, 0xed, 0xbb, 0x83,
	0xd6, 0xb1, 0x3f, 0xdc, 0x6c, 0x34, 0x2c, 0x35, 0xb9, 0x50, 0xa9, 0xd8, 0x56, 0xa0, 0x57, 0x50,
	0x17, 0xb9, 0xd4, 0xc5, 0x55, 0x5d, 0xfc, 0xf8, 0xde, 0xe2, 0x4b, 0x9d, 0x8b, 0x8b, 0x1a, 0xff,
	0x9b, 0x03, 0xbd, 0x4d, 0x6e, 0x74, 0x06, 0xdd, 0x24, 0xa5, 0x73, 0x26, 0xf2, 0xcc, 0xe4, 0x6b,
	0x31, 0xad, 0xe3, 0xfd, 0x6d, 0xea, 0xcb, 0x5c, 0x8e, 0x05, 0xe3, 0x12, 0x6f, 0x54, 0x28, 0xbd,
	0x19, 0xfd, 0x9c, 0x53, 0x1e, 0xae, 0xf4, 0x16, 0x18, 0xfd, 0x03, 0xb5, 0x2c, 0x4c, 0x59, 0x22,
	0x3d, 0xb7, 0xef, 0x0c, 0xda, 0xd8, 0x22, 0x7f, 0x0c, 0x8d, 0x82, 0x0f, 0x21, 0xa8, 0xce, 0x48,
	0x36, 0xd3, 0x9d, 0xdb, 0x58, 0x9f, 0xd1, 0x5f, 0xb0, 0xc3, 0x78, 0x44, 0x17, 0x96, 0xd0, 0x80,
	0xb5, 0x4e, 0xee, 0x7a, 0x27, 0xff, 0x14, 0xf6, 0xb6, 0xc4, 0x2b, 0x9a, 0x39, 0x89, 0x73, 0xaa,
	0xb9, 0x5d, 0x6c, 0x40, 0xe9, 0x52, 0x95, 0xb5, 0x4b, 0x7d, 0x05, 0xf4, 0x9e, 0x67, 0x09, 0xe5,
	0xb2, 0x3c, 0xe8, 0x97, 0xd0, 0x14, 0xb9, 0x0c, 0x12, 0x75, 0xd7, 0x07, 0xb8, 0xd3, 0x10, 0x85,
	0xae, 0x3f, 0xb4, 0x51, 0x71, 0x72, 0x2b, 0x72, 0x6e, 0x3c, 0x71, 0xb1, 0x45, 0xfe, 0x2f, 0x17,
	0xda, 0x57, 0xec, 0x9a, 0x33, 0x7e, 0x6d, 0x86, 0x73, 0x00, 0x4d, 0x65, 0x46, 0x20, 0x97, 0x89,
	0x51, 0xd0, 0xc1, 0x0d, 0x15, 0x98, 0x2c, 0x13, 0x5a, 0x62, 0xa9, 0x94, 0x59, 0xd0, 0x7f, 0xd0,
	0x98, 0x2e, 0x25, 0x0d, 0x3e, 0x52, 0x6a, 0xf9, 0xeb, 0x0a, 0xbf, 0xa1, 0x14, 0x1d, 0x02, 0x48,
	0x11, 0x90, 0x28, 0x4a, 0x69, 0xa6, 0x76, 0xc8, 0x19, 0x34, 0x71, 0x53, 0x8a, 0x53, 0x13, 0x40,
	0x4f, 0xa0, 0x1b, 0xce, 0x08, 0xbf, 0xa6, 0xab, 0x94, 0x1d, 0x9d, 0xd2, 0x31, 0xd1, 0x22, 0xed,
	0x7f, 0x68, 0x25, 0x29, 0x9b, 0x13, 0x49, 0x83, 0x1b, 0xba, 0xf4, 0xa0, 0xef, 0x0e, 0xda, 0x18,
	0x6c, 0xe8, 0x2d, 0x5d, 0xa2, 0x73, 0xa8, 0x1b, 0xa5, 0x99, 0xd7, 0xd2, 0x7b, 0xfa, 0x6c, 0xdb,
	0xae, 0xb2, 0xce, 0xe1, 0x95, 0xc9, 0x3e, 0xe7, 0x32, 0x5d, 0xe2, 0xa2, 0x16, 0x9d, 0x40, 0x35,
	0x97, 0x0b, 0xe1, 0xb5, 0x35, 0xc7, 0xd1, 0x36, 0xc7, 0xf6, 0xac, 0xb0, 0xae, 0x40, 0x47, 0xd0,
	0xcd, 0x33, 0x1a, 0xdc, 0x92, 0x45, 0x60, 0x2d, 0xea, 0xf4, 0x9d, 0x41, 0x03, 0xb7, 0xf3, 0x8c,
	0xbe, 0x23, 0x8b, 0x53, 0x63, 0xd4, 0x01, 0x34, 0x15, 0x99, 0x71, 0xb7, 0x6b, 0xdc, 0x55, 0x01,
	0xed, 0xee, 0x0b, 0xa8, 0x26, 0x31, 0xe1, 0xde, 0xae, 0x9e, 0xf7, 0xa3, 0x7b, 0x1f, 0xda, 0x38,
	0x26, 0x1c, 0xeb, 0xf4, 0xfd, 0x11, 0xb4, 0xcb, 0x62, 0x50, 0x0f, 0x5c, 0xe5, 0x91, 0xa3, 0x7d,
	0x54, 0xc7, 0xbb, 0x8d, 0x34, 0x3b, 0x61, 0xc0, 0xa8, 0x72, 0xe2, 0xf8, 0x3f, 0x1d, 0xd8, 0xdd,
	0x60, 0x2d, 0x0d, 0xd9, 0x59, 0x1b, 0xf2, 0x53, 0xe8, 0x91, 0x39, 0x61, 0x31, 0x99, 0xc6, 0x34,
	0x58, 0x5b, 0x83, 0xdd, 0x55, 0xdc, 0xca, 0xec, 0x81, 0x7b, 0xb7, 0x0a, 0xea, 0xa8, 0x48, 0xcd,
	0x44, 0xf5, 0x0a, 0xb8, 0xd8, 0x22, 0x34, 0x82, 0x1d, 0x65, 0x9f, 0x1a, 0xfb, 0xc3, 0x1d, 0x37,
	0x25, 0xca, 0xcc, 0x69, 0x4a, 0x78, 0x38, 0x0b, 0x58, 0xe4, 0xd5, 0xb4, 0xb4, 0x86, 0x09, 0x5c,
	0x44, 0xfe, 0x77, 0x07, 0x3a, 0x76, 0xe0, 0xf6, 0x5d, 0xbe, 0x86, 0x96, 0xbc, 0x23, 0xb1, 0xaf,
	0xea, 0xf0, 0x5e, 0x97, 0x71, 0xb9, 0x42, 0xfd, 0xfb, 0x52, 0x1e, 0x8a, 0x88, 0x46, 0xd6, 0xc8,
	0x02, 0xaa, 0x2d, 0x2e, 0x25, 0xaa, 0xeb, 0xb8, 0x66, 0x8b, 0x4b, 0xd1, 0x8b, 0x48, 0xcd, 0x81,
	0xa6, 0xa9, 0x48, 0xed, 0x33, 0x30, 0xe0, 0xec, 0x5f, 0xf8, 0xfb, 0x0b, 0x89, 0x63, 0x2a, 0x87,
	0xa1, 0x48, 0xe9, 0xf0, 0x13, 0x67, 0xe6, 0x7b, 0x30, 0xad, 0xe9, 0x9f, 0xe7, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x53, 0xd6, 0xbb, 0x06, 0x27, 0x06, 0x00, 0x00,
}
