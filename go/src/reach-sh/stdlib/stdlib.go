package stdlib

import (
  "github.com/ethereum/go-ethereum/common"
  "math/big"

  "github.com/ethereum/go-ethereum/common/math"
  solsha3 "github.com/miguelmota/go-solidity-sha3"
)

var (
  bigZero = new(big.Int)
  tt255   = math.BigPow(2, 255)
  tt256   = math.BigPow(2, 256)
  tt256m1 = new(big.Int).Sub(tt256, big.NewInt(1))
)

type Type_bool = bool
type Type_uint256 = *big.Int
type Type_address = common.Address
type Type_uint256arr = []Type_uint256

type Txn struct { 
     Balance Type_uint256
     Value Type_uint256
     DidTimeout bool
     From Type_address
     Data Msg }

var Txn0 = Txn{ Balance: big.NewInt(0), Value: big.NewInt(0) }

func Ite_uint256(c Type_bool, x Type_uint256, y Type_uint256) Type_uint256 {
  if c {
    return x
  } else {
    return y } }

func Ite_bool(c Type_bool, x Type_bool, y Type_bool) Type_bool {
  if c {
    return x
  } else {
    return y } }

func Eq(x Type_uint256, y Type_uint256) Type_bool {
  return x.Cmp(y) == 0 }

func Add(x Type_uint256, y Type_uint256) Type_uint256 {
  return math.U256(new(big.Int).Add(x, y))
}
func Sub(x Type_uint256, y Type_uint256) Type_uint256 {
  return math.U256(new(big.Int).Sub(new(big.Int).Add(x, tt256), y)) }

// This currently panics if y == 0. What behavior should it have?
func Div(x Type_uint256, y Type_uint256) Type_uint256 {
  return new(big.Int).Div(x, y) }

// This currently panics if y == 0. What behavior should it have?
func Mod(x Type_uint256, y Type_uint256) Type_uint256 {
  return new(big.Int).Mod(x, y) }

func Gt(x Type_uint256, y Type_uint256) Type_bool {
  return x.Cmp(y) > 0 }

func Lt(x Type_uint256, y Type_uint256) Type_bool {
  return x.Cmp(y) < 0 }

func Keccak256(x Type_uint256, y Type_uint256) Type_uint256 {
  panic("XXX") }

func Assert( b Type_bool ) {
  if b {
    return
  } else {
    panic("Assertion failed") } }

type Msg int

var Msg0 Msg = 0

func MsgEncode_uint256( m Msg, v Type_uint256 ) Msg {
  panic("XXX") }

func MsgEncode_bool( m Msg, v Type_bool ) Msg {
  panic("XXX") }

func MsgEncode_uint256arr( m Msg, v []Type_uint256 ) Msg {
  panic("XXX") }

func MsgEncode_address( m Msg, v Type_address ) Msg {
  panic("XXX") }

func MsgDecode_bool( m Msg, path []string ) Type_bool {
  panic("XXX") }

func MsgDecode_uint256( m Msg, path []string ) Type_uint256 {
  panic("XXX") }

func MsgDecode_uint256arr( m Msg, path []string ) []Type_uint256 {
  panic("XXX") }

type Contract int

func (c Contract) SendRecv(dbg string, m_name string, m Msg, amount Type_uint256, e_name string, delay Type_uint256, t_name string ) <-chan Txn {
  panic("XXX SendRecv")
}

func (c Contract) Recv(dbg string, e_name string, delay Type_uint256, to_me bool, to_m Msg, to_m_name string, to_e_name string ) <-chan Txn {
  panic("XXX Recv")
}
