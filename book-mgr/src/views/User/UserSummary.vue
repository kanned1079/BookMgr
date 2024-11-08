<script setup lang="ts">
import {onMounted, ref} from 'vue'
import {useMessage} from 'naive-ui'
import {useUserStore} from "@/stores/userinfo";
import instance from "@/axios";

const userStore = useUserStore()
const message = useMessage()
let animated = ref<boolean>(false)

interface UserSummary {
  unreturned: number
  borrowed_nums: number
  ranking_percent: number
}

let summary = ref<UserSummary>({
  unreturned: 0,
  borrowed_nums: 0,
  ranking_percent: 0.00,
})

let getUserSummary = async () => {
  try {
    let {data} = await instance.get('/api/user/v1/summary', {
      params: {
        user_id: userStore.thisUser.id || -1
      }
    })
    if (data.code === 200) {
      Object.assign(summary.value, data.summary)
    }
  } catch (err: any) {
    message.error(err + '')
  }
}


onMounted(async () => {
  await getUserSummary()
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'UserSummary'
}
</script>

<template>
  <div style="padding: 20px 20px 0 20px;">
    <n-card hoverable :embedded="true" :bordered="false">
      <div style="display: flex; flex-direction: row; justify-content: space-between; align-items: center">
        <p style="font-size: 1.1rem; font-weight: 500">用户总览</p>
      </div>
    </n-card>
  </div>

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card hoverable :embedded="true" :bordered="false">
        <n-statistic label="未归还书目数量" tabular-nums>
          <n-number-animation ref="numberAnimationInstRef" :from="0" :to="summary.unreturned"/>
          <template #suffix>
            本
          </template>
        </n-statistic>
      </n-card>

      <n-card style="margin-top: 20px;" hoverable :embedded="true" :bordered="false">
        <div style="display: flex; flex-direction: row; justify-content: left; align-items: center;">
          <n-statistic label="总借阅量 / 在所有用户的百分比" tabular-nums>
            <n-number-animation ref="numberAnimationInstRef" :from="0" :to="summary.borrowed_nums"/>
            本&nbsp;/&nbsp;
            <n-number-animation ref="numberAnimationInstRef" :from="0" :to="summary.ranking_percent" :precision="2"/>
            %
          </n-statistic>
        </div>
      </n-card>
    </div>
  </transition>
</template>

<style scoped lang="less">
.root {
  padding: 20px;
}

</style>