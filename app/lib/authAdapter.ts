
import type { Adapter, AdapterAccount, AdapterUser } from "next-auth/adapters"

export function mapExpiresAt(account: AdapterAccount): AdapterAccount {
  const expires_at: number = parseInt(account.expires_at as unknown as string)
  return {
    ...account,
    expires_at,
  }
}



export default function MyAdapter(): Adapter {

    const apiBase = 'http://localhost:8080/auth'

    return {
      async createUser(user) {
        const response = await fetch(`${apiBase}/user`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(user),
            })
        const data = await response.json()
        return data

      },
      async getUser(id) {
 
        const response = await fetch(`${apiBase}/user/${id}`)
        const data = await response.json()
        return data
        
      },
      async getUserByEmail(email) {

        const response = await fetch(`${apiBase}/user/email/${email}`, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' }
        })

        if (response.ok) {
            const data = await response.json()
      
            return data

        } else {
            const data = await response.text()
            console.log("email data", data)
            return null
        }
    

        
      },
      async getUserByAccount({ providerAccountId, provider }) {
    
        const response = await fetch(`${apiBase}/user/account/${provider}/${providerAccountId}`)
        const data = await response.json()
        return data

      },

      async updateUser(user) {
    
        const response = await fetch(`${apiBase}/user`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(user),
            })

            if (!response.ok) {
              const data = await response.text()
              console.log("updateUser error", data)
              return null
            }
        const data = await response.json()
        
        return data
      },

      async deleteUser(userId) {
        const response = await fetch(`${apiBase}/user/${userId}`, {
            method: 'DELETE',
            })
        const data = await response.json()
        return data
      },
      async linkAccount(account) {

        const response = await fetch(`${apiBase}/link`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(account),
            })
        const data = await response.json()

        const res = mapExpiresAt(data)
    
        return res
      },
      async unlinkAccount({ providerAccountId, provider }) {
        const response = await fetch(`${apiBase}/link`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ providerAccountId, provider }),
            })
            
        const data = await response.json()
        return data
      },
      async createSession({ sessionToken, userId, expires }) {
        console.log("createSession", sessionToken, userId, expires)
        const response = await fetch(`${apiBase}/session`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ sessionToken, userId, expires }),
            })
        console.log("createSession response", response)
        if (!response.ok) {
          const data = await response.text()
          console.log("createSession error", data)
          return null
        }
        const data = await response.json()

        const res = {
          ...data,
          expires: new Date(data.expires)
        }

        console.log("createSession res", res)
   
        return res
      },
      async getSessionAndUser(sessionToken) {
       
        const response = await fetch(`${apiBase}/session/${sessionToken}`)
        
        if (!response.ok) {
          const data = await response.text()
          console.log("getSessionAndUser error", data)
          return null
        }
        
        const data = await response.json()
        const res = {
          ...data,
          session: { 
            ...data.session,
            expires: new Date(data.session.expires)
          }
        }
        return res
      },
      async updateSession({ sessionToken }) {
      
        const response = await fetch(`${apiBase}/session`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ sessionToken }),
            })

        const data = await response.json()
        return data
      },
      async deleteSession(sessionToken) {
        
        const response = await fetch(`${apiBase}/session/${sessionToken}`, {
            method: 'DELETE',
            })
            if (!response.ok) {
              const data = await response.text()
              console.log("deleteSession error", data)
              return null
            }
        const data = await response.json()

        return data
      },
      async createVerificationToken({ identifier, expires, token }) {
        
        const response = await fetch(`${apiBase}/verification-token`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ identifier, expires, token }),
            })
        const data = await response.json()
        return data
      },
      async useVerificationToken({ identifier, token }) {
        
        const response = await fetch(`${apiBase}/verification-token`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ identifier, token }),
            })

            if (!response.ok) {
              const data = await response.text()
              console.log("useVerificationToken error", data)
              return null
            }

        const data = await response.json()

        return data
      },
    }
  }