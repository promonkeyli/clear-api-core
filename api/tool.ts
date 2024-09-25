import request from '@/utils/http;' 

/** 获取工具 */
export async function GetTools(options?: { [key: string]: any}){
  return request<any>('/tools', {
       method: 'GET',
       ...(options || {})
  })
}

/** 添加工具 */
export async function PostTools( body: API.model.Tool, options?: { [key: string]: any}){
  return request<any>('/tools', {
       method: 'POST',
       header: {
         'Content-Type': 'application/json'
       },
       body,
       ...(options || {})
  })
}
