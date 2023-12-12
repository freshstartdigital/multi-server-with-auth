

export class Config {
    apiUrl: string

    constructor() {
        this.apiUrl = process.env.NODE_ENV == 'production' ? 'http://api:8080' : 'http://localhost:8080'
    }
}