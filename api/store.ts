import request from '@/utils/http;' 


/** For valid response try integer IDs with positive integer value. Negative or non-integer values will generate API errors */
export async function deleteOrder( params: number,options?: { [key: string]: any}){
  return request<any>(`/store/order/${orderId}`, {
       method: 'DELETE',
       ...(options || {})
  })
}

/** For valid response try integer IDs with value >= 1 and <= 10. Other values will generated exceptions */
export async function getOrderById( params: number,options?: { [key: string]: any}){
  return request<any>(`/store/order/${orderId}`, {
       method: 'GET',
       ...(options || {})
  })
}

export async function placeOrder( body: API.Order,options?: { [key: string]: any}){
  return request<any>('/store/order', {
       method: 'POST',
       header: {
         'Content-Type': 'application/json'
       },
       body,
       ...(options || {})
  })
}

/** Returns a map of status codes to quantities */
export async function getInventory(options?: { [key: string]: any}){
  return request<any>('/store/inventory', {
       method: 'GET',
       ...(options || {})
  })
}
