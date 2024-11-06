import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

interface User {
    id: number
    role: string
    email: string
    created_at: string
}

export const useUserStore = defineStore('user', () => {
    // const count = ref(0)
    // const doubleCount = computed(() => count.value * 2)
    // function increment() {
    //     count.value++
    // }

    // return { count, doubleCount, increment }

    let authed = ref<boolean>(false)

    let thisUser = ref<User>({
        id: 1,
        role: 'user',
        email: 'kanned1079@gmail.com',
        created_at: '2009-10-12'
    })

    return {
        authed,
        thisUser,
    }
})
