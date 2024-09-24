import request from '@/utils/http;' 


/** Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing. */
export async function findPetsByTags(options?: { [key: string]: any}){
  return request<any>('/pet/findByTags', {
       method: 'GET',
       ...(options || {})
  })
}

/** Multiple status values can be provided with comma separated strings */
export async function findPetsByStatus(options?: { [key: string]: any}){
  return request<any>('/pet/findByStatus', {
       method: 'GET',
       ...(options || {})
  })
}

export async function uploadFile( params: number, body: { additionalMetadata: string; file: string; },options?: { [key: string]: any}){
  return request<any>(`/pet/${petId}/uploadImage`, {
       method: 'POST',
       header: {
         'Content-Type': 'multipart/form-data'
       },
       body,
       ...(options || {})
  })
}

export async function addPet( body: API.Pet,options?: { [key: string]: any}){
  return request<any>('/pet', {
       method: 'POST',
       header: {
         'Content-Type': 'application/json'
       },
       body,
       ...(options || {})
  })
}

export async function updatePet( body: API.Pet,options?: { [key: string]: any}){
  return request<any>('/pet', {
       method: 'PUT',
       header: {
         'Content-Type': 'application/xml'
       },
       body,
       ...(options || {})
  })
}

export async function deletePet( params: number,options?: { [key: string]: any}){
  return request<any>(`/pet/${petId}`, {
       method: 'DELETE',
       ...(options || {})
  })
}

/** Returns a single pet */
export async function getPetById( params: number,options?: { [key: string]: any}){
  return request<any>(`/pet/${petId}`, {
       method: 'GET',
       ...(options || {})
  })
}

export async function updatePetWithForm( params: number, body: { name: string; status: string; },options?: { [key: string]: any}){
  return request<any>(`/pet/${petId}`, {
       method: 'POST',
       header: {
         'Content-Type': 'application/x-www-form-urlencoded'
       },
       body,
       ...(options || {})
  })
}
