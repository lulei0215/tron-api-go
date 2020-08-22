package tron

import ()

type TronKit struct{
  tronApi *TronApi
  credential *Credential
}

func NewTronKit(tronApi *TronApi, credential *Credential) *TronKit{
  return &TronKit{
    tronApi: tronApi,
    credential: credential,
  }
}

func (kit *TronKit) GetDefaultAddress() string {
  return kit.credential.AddressBase58()
}

func (kit *TronKit) SendTrx(to string, amount int64) (*TransactionResult, error){
  tx, err := kit.tronApi.CreateTransaction(to, amount, kit.credential.AddressBase58())
  if err != nil {
    return nil, err
  }
  
  sig, err := kit.credential.SignHex(tx.TxId)
  if err != nil {
    return nil, err
  }
  tx.Signature = []string{ sig }
  
  succeed, err := kit.tronApi.BroadcastTransaction(tx)
  if err != nil {
    return nil, err
  }
  return &TransactionResult{tx.TxId, succeed}, nil
  
}

func (kit *TronKit) GetTrxBalance(address string) (int64,error){
  account, err := kit.tronApi.GetAccount(address)
  if err != nil {
    return 0, err
  }  
  return account.Balance, nil
}

func (kit *TronKit) Trc20(contractAddress string) (*Trc20, error) {
  return NewTrc20(kit.tronApi, kit.credential, contractAddress)
}