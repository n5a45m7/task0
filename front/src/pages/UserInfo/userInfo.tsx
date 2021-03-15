import { observer } from 'mobx-react';
import React, { useEffect } from 'react';
import { Account } from '../../components/account';
import { UserInfo } from '../../components/user';
import { IUser, IUserStore } from '../../stores';

interface IProps {
  store: IUserStore
  user: IUser
}

// export const PageUserInfo = observer(({id, name, surname, load}: IProps) => {
const PageUserInfoComponent = ({store, user}: IProps) => {
  const {id, name, surname, accounts, load, loading, error} = store
  useEffect(() => {
    console.log('use effect call')
    load(user)
  }, [load, user]);
  return (
    <div>
        {loading 
          ? <div>Loading...</div>
          : <div>
            {error
              ? error
              : 
              <>
                <div className="yellow">User Info:</div>
                <UserInfo id={id} name={name} surname={surname}></UserInfo>           
                { accounts?.length > 0 && <>
                  <div className="yellow">Accounts:</div>
                  <div>
                    {accounts.map((acc, index) => 
                      <div key={index}><Account id={acc.id} balance={acc.balance} txs={acc.txs}></Account></div>
                    )}
                  </div>
                </> }

              </>
            }
            </div>
        }
    </div>
  );
}

export const PageUserInfo = observer(PageUserInfoComponent)