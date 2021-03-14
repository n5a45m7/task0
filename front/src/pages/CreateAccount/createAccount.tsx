import { observer } from 'mobx-react';
import React, { useCallback, useEffect, useState } from 'react';
import { ICreateAccountStore, IUser } from '../../stores';

interface IProps {
  store: ICreateAccountStore
  user: IUser
}

const PageCreateAccountComponent = ({store, user}: IProps) => {
  const {response, error, loading, create, reset} = store

  useEffect(() => {
    return () => {
      reset()
    }
  }, [reset])

  const [initalCredit, setInitialCredit] = useState(0)

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const val = +event.currentTarget.value
    if (val >= 0) {
      setInitialCredit(val)
    }
  }

  const handleClick = useCallback(() => {
    create({customerID: user.id, initialCredit: initalCredit})
  }, [create, user, initalCredit])

  const createAccountComponent = (
    <div>
      <div>Create Account</div>
      <p></p>
      <div>
        <label>
          Inital Credit:
          <input type="number" min={0} value={initalCredit} onChange={handleChange} />
        </label>
      </div>
      <div>
        <button onClick={handleClick}> Create </button>
      </div>
    </div>
  )

  return (
    <div>
        {loading 
          ? <div>Loading...</div>
          : <div>
            {error
              ? error
              : response?.id
                ?
                  <div>
                    Created account with id = {response?.id}
                    <a onClick={() => { reset() }}>Create one more</a>
                  </div>
                :
                createAccountComponent
            }
            </div>
        }
    </div>
  );
}

export const PageCreateAccount = observer(PageCreateAccountComponent)