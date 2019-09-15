<template>
    <div>
        <input v-model="lastName" placeholder="姓">
        <input v-model="firstName" placeholder="名">
        <button @click="addUserFromForm" class="button">登録</button>
        <div class="table">
            <div class="table-row">
                <div class="table-column table-column-1">ID</div><div class="table-column table-column-2">名前</div>
            </div>
            <template v-for="user in users">
                <div :key="user.ID" class="table-row">
                    <div class="table-column table-column-1">{{user.ID}}</div><div class="table-column table-column-2">{{user.LastName}} {{user.FirstName}}</div>
                </div>
            </template>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import { mapState, mapActions } from 'vuex';

export default {
    created() {
        this.getAllUsers()
    },
    data() {
        return {
            lastName: "",
            firstName: "",
        }
    },
    computed: mapState({
        users: 'users'
    }),
    methods: {
        ...mapActions([
            'getAllUsers',
            'addUser',
        ]),
        async addUserFromForm() {
            await this.addUser({
                firstName: this.firstName,
                lastName: this.lastName,
            })
            this.firstName = "";
            this.lastName = "";
            this.getAllUsers()
        },
    }
}
</script>

<style lang="scss">
.table {
    margin: 10px;
    .table-row {
        display: flex;
        align-items: center;
        height: 30px;
    }
    .table-column {
        text-align: center;
        &.table-column-1 {
            width: 50px;
        }
        &.table-column-2 {
            width: 150px;
        }
    }
}
.button {
    margin: 10px;
}
</style>
