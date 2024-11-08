<script setup lang="ts">
import {ref, onMounted} from 'vue'
import {useMessage} from 'naive-ui'
import instance from "@/axios";

const message = useMessage()
let animated = ref<boolean>(false)

interface AdminSummary {
  user_count: number
  book_count: number
  borrowed_count: number
}

let summary = ref<AdminSummary>({
  user_count: 0,
  book_count: 0,
  borrowed_count: 0,
})

let getAdminSummary = async () => {
  try {
    let {data} = await instance.get('/api/admin/v1/summary')
    if (data.code === 200) {
      Object.assign(summary.value, data.summary)
    } else {
      message.error('未知错误' + data.msg || '')
    }
  } catch (err: any) {
    message.error(err + '')
  }
}

onMounted(async () => {
  await getAdminSummary()
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'AdminSummary'
}
</script>

<template>
  <div style="padding: 20px 20px 0 20px;">
    <n-card hoverable :embedded="true" :bordered="false">
      <div style="display: flex; flex-direction: row; justify-content: space-between; align-items: center">
        <p style="font-size: 1.1rem; font-weight: 500">总览</p>
      </div>
    </n-card>
  </div>

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card hoverable :embedded="true" :bordered="false">
        <n-statistic label="图书馆总注册人数" tabular-nums>
          <n-number-animation ref="numberAnimationInstRef" :from="0" :to="summary.user_count" />
          <template #suffix>
            人
          </template>
        </n-statistic>
      </n-card>

      <n-card style="margin-top: 20px" hoverable :embedded="true" :bordered="false">
        <n-statistic label="总藏书量" tabular-nums>
          <n-number-animation ref="numberAnimationInstRef" :from="0" :to="summary.book_count" />
          <template #suffix>
            本
          </template>
        </n-statistic>
      </n-card>

      <n-card style="margin-top: 20px" hoverable :embedded="true" :bordered="false">
        <n-statistic label="总借阅量" tabular-nums>
          <n-number-animation ref="numberAnimationInstRef" :from="0" :to="summary.borrowed_count" />
          <template #suffix>
            本书
          </template>
        </n-statistic>
      </n-card>
    </div>
  </transition>


</template>

<style scoped lang="less">
.root {
  padding: 20px;
}

</style>