import { HomePageLayoutProps } from "@/types/Layout";
import {  GetServerSidePropsContext, GetServerSidePropsResult } from "next";
import { Config } from "./config";


export class Ssr extends Config {
    private context: GetServerSidePropsContext

    constructor(context: GetServerSidePropsContext) {
        super()
        this.context = context
    }

    async getHomePageLayoutProps(): Promise<GetServerSidePropsResult<HomePageLayoutProps>> {
        const data = await fetch(`${this.apiUrl}/hello`)
        const res = await data.json()
        return {
          props: {
            text: res.text,
          },
        }
    }
}