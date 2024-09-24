import request from '@/utils/http;' 


export async function createUsersWithListInput( body: Array<unknown>,options?: { [key: string]: any}){
  return request<any>('/user/createWithList', {
       method: 'POST',
       header: {
         'Content-Type': 'application/json'
       },
       body,
       ...(options || {})
  })
}

/** This can only be done by the logged in user. */
export async function createUser( body: API.User,options?: { [key: string]: any}){
  return request<any>('/user', {
       method: 'POST',
       header: {
         'Content-Type': 'application/json'
       },
       body,
       ...(options || {})
  })
}

export async function logoutUser(options?: { [key: string]: any}){
  return request<any>('/user/logout', {
       method: 'GET',
       ...(options || {})
  })
}

export async function loginUser(options?: { [key: string]: any}){
  return request<any>('/user/login', {
       method: 'GET',
       ...(options || {})
  })
}

export async function createUsersWithArrayInput( body: Array<unknown>,options?: { [key: string]: any}){
  return request<any>('/user/createWithArray', {
       method: 'POST',
       header: {
         'Content-Type': 'application/json'
       },
       body,
       ...(options || {})
  })
}

export async function getUserByName( params: string,options?: { [key: string]: any}){
  return request<any>(`/user/${username}`, {
       method: 'GET',
       ...(options || {})
  })
}

/** This can only be done by the logged in user. */
export async function updateUser( params: string, body: API.User,options?: { [key: string]: any}){
  return request<any>(`/user/${username}`, {
       method: 'PUT',
       header: {
         'Content-Type': 'application/json'
       },
       body,
       ...(options || {})
  })
}

/** This can only be done by the logged in user. */
export async function deleteUser( params: string,options?: { [key: string]: any}){
  return request<any>(`/user/${username}`, {
       method: 'DELETE',
       ...(options || {})
  })
}
