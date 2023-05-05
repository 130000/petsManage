const env = import.meta.env.MODE || 'prod'
const envConfig = {
    development: {
        baseApi: 'http://localhost:8080',
        mockApi: '/home'
    },
    test: {
        baseApi: '',
        mockApi: '/home'
    },
    prod: {
        baseApi: '',
        mockApi: '/home'
    }
}
export default {
    env,
    mock: false,
    ...envConfig[env]
}
