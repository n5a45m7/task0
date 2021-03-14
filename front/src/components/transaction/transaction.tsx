import { observer } from 'mobx-react';
import React from 'react';

interface IProps {
  id: number
  amount: number
}

export const TransactionComponent = ({id, amount}: IProps) => {
  return (
    <div>
        <div>ID: {id}</div>
        <div>Amount: {amount}</div>
    </div>
  );
}

export const Transaction = observer(TransactionComponent)