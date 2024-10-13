import {redirect} from '@sveltejs/kit'

export const actions = {
  default: async ({cookie, request}) => {
    const form = await request.formData()

    const response = await fetch("http://localhost:8000/api/v1/register", {
      method: "POST",
      headers: {
      "Content-Type": "application/json"
      },
      body: JSON.stringify({
        name: form.get('username'),
        email: form.get('email'),
        password: form.get('password'),
        password_confirm: form.get('password_confirmation')
      })
      })
  
    const data = await response.json()
  
    if(response.ok && response.status == 201) {
      throw redirect(303, '/login')
    }

    return data
  }
}