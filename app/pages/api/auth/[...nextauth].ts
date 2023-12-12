import MyAdapter from "@/lib/authAdapter"
import NextAuth from "next-auth"
import EmailProvider from "next-auth/providers/email"

export const authOptions = {
  providers: [
    EmailProvider({
        server: {
            host: "smtp.sendgrid.net",
            port: 587,
            auth: {
                user: "apikey", // This is literally the string "apikey"
                pass: process.env.SENDGRID_API_KEY
            }
        },
        from: 'dev@droneanalytics.com.au'
    }),
    // ... other providers
],
  adapter: MyAdapter(),
}

export default NextAuth(authOptions)