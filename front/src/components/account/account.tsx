import { observer } from 'mobx-react';
import React from 'react';
import { Transaction } from '../transaction';
interface IProps {
  id: number
  balance: number
  txs: {
      id: number
      amount: number
  }[]
}

export const AccountComponent = ({id, balance, txs}: IProps) => {
  return (
    <div>
        <div>ID: {id}</div>
        <div>Balance: {balance}</div>
        {txs && <>
            <div>Transactions:</div>
            {txs.map((tx, index) => <div key={index}>
                <Transaction id={tx.id} amount={tx.amount}></Transaction>
            </div>)}
        </>}
    </div>
  );
}

export const Account = observer(AccountComponent)