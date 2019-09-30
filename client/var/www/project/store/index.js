import axios from 'axios';

export const state = () => ({
    users: []
})

export const mutations = {
    setUsers(state, users) {
        state.users = users;
    }
}

export const actions = {
    async getAllUsers(context) {
        try {
            const { data } = await axios.get('/api/users')
            context.commit('setUsers', data)
        } catch (error) {
            console.log(error)
        }
    },
    async addUser(context, params) {
        try {
            await axios.post('/api/users', params)
        } catch (error) {
            console.log(error)
        }
    }
}
