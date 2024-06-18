'use client';

import { Bond, Pagination } from '../types';
import { API_URL } from '../utils/constants';
import { ErrorWithRequest } from './Error';

export async function userLogin(credentials: Record<string, unknown>) {
  const res = await fetch(API_URL + '/auth/login', {
    method: 'POST',
    mode: 'cors',
    body: JSON.stringify(credentials),
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  });
  if (!res.ok) {
    const json = await res.json();
    let message = '';
    if (typeof json.error === 'string') {
      message = json.error;
    }
    if (typeof json.error === 'object') {
      const { error } = json;
      const keys = Object.keys(error)[0];
      message = error[keys];
    }
    throw new ErrorWithRequest(message, res.url);
  }
  return await res.json();
}

export async function validateSession() {
  const res = await fetch(API_URL + '/auth/session', {
    method: 'GET',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
    },
  });
  return await res.json();
}

export async function userLogout(data: any) {
  const res = await fetch(API_URL + '/auth/logout', {
    method: 'POST',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
    },
  });
  return await res.json();
}

export async function userRegister(credentials: Record<string, unknown>) {
  const res = await fetch(API_URL + '/auth/join', {
    method: 'POST',
    mode: 'cors',
    body: JSON.stringify(credentials),
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  });
  if (!res.ok) {
    const json = await res.json();
    let message = '';
    if (typeof json.error === 'string') {
      message = json.error;
    }
    if (typeof json.error === 'object') {
      const { error } = json;
      const keys = Object.keys(error)[0];
      message = error[keys];
    }
    throw new ErrorWithRequest(message, res.url);
  }
  return await res.json();
}

export async function getUserOwnedBonds({queryKey}:{queryKey: [string, number]}) {
  console.log(queryKey)
  const page = queryKey[1]
  const params = new URLSearchParams();
  params.append('page', page.toString());

  const res = await fetch(API_URL + '/bonds?'+params.toString(), {
    method: 'GET',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',

    },
  })
  
  const json = await res.json();
  if (!res.ok) {
    const json = await res.json();
    const message = handleError(json);
    throw new ErrorWithRequest(message, res.url);
  }
  return json as { bonds: Bond[], pagination: Pagination };
}


const handleError = function (json:Record<string, unknown>) {
  if (!json) {
    return '';
  }
    let message = '';
    if (typeof json.error === 'string') {
      message = json.error;
    }
    if (typeof json.error === 'object') {
      const { error } = json;
      const keys = Object.keys(error as object)[0];
     if (!error) return 'A error occurred';
      message = error[keys as keyof typeof error];
    }

    return message;
}