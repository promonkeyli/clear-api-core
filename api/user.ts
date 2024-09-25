import request from '@/utils/http;' 

/** 使用用户名密码进行登录 */
export async function PostLogin( body: API.model.User, options?: { [key: string]: any}){
  return request<any>('/login', {
       method: 'POST',
       header: {
         'Content-Type': 'application/json'
       },
       body,
       ...(options || {})
  })
}

/** 用户注销，清除会话、注销令牌 */
export async function PostLogout(options?: { [key: string]: any}){
  return request<any>('/logout', {
       method: 'POST',
       ...(options || {})
  })
}

/** 获取所有用户 */
export async function GetUsers(options?: { [key: string]: any}){
  return request<any>('/users', {
       method: 'GET',
       ...(options || {})
  })
}

/** 新增用户 */
export async function PostUsers( body: API.model.User, options?: { [key: string]: any}){
  return request<any>('/users', {
       method: 'POST',
       header: {
         'Content-Type': 'application/json'
       },
       body,
       ...(options || {})
  })
}
