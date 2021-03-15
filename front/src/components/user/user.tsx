import { observer } from 'mobx-react';
import React from 'react';

interface IProps {
  id: number
  name: string
  surname: string
}

const UserInfoComponent = ({id, name, surname}: IProps) => {
  return (
    <div className="userInfo">
        <div>ID: {id}</div>
        <div>Name: {name}</div>
        <div>Surname: {surname}</div>
    </div>
  );
}

export const UserInfo = observer(UserInfoComponent)