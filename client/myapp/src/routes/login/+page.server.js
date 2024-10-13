import { redirect } from "@sveltejs/kit"

export const actions = {
  default: async ({cookies, request}) => {
    const form = await request.formData()

    const response = await fetch("http://localhost:8000/api/v1/login", {
      method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          email: form.get('email'),
          password: form.get('password')
        }),
        credentials: "include"
    })

    const json = await data.json()

    if(response.ok && response.status == 201) {
      localStorage.setItem('token', data.token)
      redirect(303, '/')
    }

    return json
  }
}
